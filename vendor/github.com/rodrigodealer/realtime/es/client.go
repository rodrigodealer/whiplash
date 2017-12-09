package es

import (
	"context"
	"errors"
	"log"
	"os"
	"reflect"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/rodrigodealer/realtime/models"
	"github.com/rodrigodealer/realtime/tracing"

	elastic "gopkg.in/olivere/elastic.v5"
)

type ElasticSearch interface {
	Connect()
	Ping() int
	IndexUser(index string, user *models.FacebookUser, span *tracing.Tracing)
	GetUser(index string, ID string) (models.FacebookUser, error)
}

type EsClient struct {
	Client *elastic.Client
}

func (e *EsClient) Connect() {
	var url = config(os.Getenv("ELASTICSEARCH_URL"))
	client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false))
	if err != nil {
		log.Panicf("Error trying to connect to ElasticSearch: \n %s \n %s", url, err)
	} else {
		e.Client = client
	}
}

func (e *EsClient) Ping() int {
	if e.Client != nil {
		ctx := context.Background()
		var url = config(os.Getenv("ELASTICSEARCH_URL"))
		info, code, _ := e.Client.Ping(url).Do(ctx)
		log.Printf("Elasticsearch returned with code %d and version %s", code, info.Version.Number)
		return code
	} else {
		return 500
	}
}

func (e *EsClient) IndexSetup(index string) {
	ctx := context.Background()
	exists, err := e.Client.IndexExists(index).Do(ctx)
	if err != nil {
		log.Panic(err)
	}
	if !exists {
		_, err := e.Client.CreateIndex(index).Do(ctx)
		if err != nil {
			log.Panic(err)
		} else {
			log.Printf("Index %s created", index)
		}
	} else {
		log.Printf("Index %s already exists, skipping", index)
	}
}

func (e *EsClient) IndexUser(index string, user *models.FacebookUser, parentSpan *tracing.Tracing) {
	if parentSpan.Span != nil {
		span := opentracing.StartSpan("Index user", opentracing.ChildOf(parentSpan.Span.Context()))
		defer span.Finish()
	}
	ctx := context.Background()
	indexed, err := e.Client.Index().
		Index(index).
		Type(index).
		Id(user.ID).
		BodyJson(user).
		Do(ctx)
	if err != nil {
		log.Printf("Error indexing user: %s \n %s", user.ToJson, err.Error)
	}
	log.Printf("Indexed tweet %s to index %s\n", indexed.Id, index)
}

func (e *EsClient) GetUser(index string, ID string) (models.FacebookUser, error) {

	ctx := context.Background()
	termQuery := elastic.NewTermQuery("id", ID)
	searchResult, err := e.Client.Search().
		Index(index).
		Query(termQuery).
		From(0).Size(10).
		Pretty(true).
		Do(ctx)
	if err != nil {
		log.Printf("Error searching for user: \n %s", err)
	}

	log.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
	var facebookUser models.FacebookUser
	if searchResult.Hits.TotalHits > 0 {
		log.Printf("Found a total of %d users\n", searchResult.TotalHits())

		for _, item := range searchResult.Each(reflect.TypeOf(facebookUser)) {
			if user, ok := item.(models.FacebookUser); ok {
				log.Printf("User %s: %s\n", user.ID, user.Name)
				return user, nil
			}
		}
	} else {
		log.Printf("No user found: %s", ID)
	}
	return facebookUser, errors.New("No user found")
}
