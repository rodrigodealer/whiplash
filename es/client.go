package es

import (
	"context"
	"log"
	"time"

	"github.com/rodrigodealer/whiplash/util"
	elastic "gopkg.in/olivere/elastic.v5"
)

type ElasticSearch interface {
	Connect()
	Ping() int
}

type EsClient struct {
	Client *elastic.Client
}

func (e *EsClient) Connect() {
	var url = util.EnvOrElse("ELASTICSEARCH_URL", "http://127.0.0.1:9200")
	client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false),
		elastic.SetHealthcheckTimeoutStartup(200*time.Millisecond))
	if err != nil {
		log.Panicf("Error trying to connect to ElasticSearch: \n %s \n %s", url, err)
	} else {
		e.Client = client
	}
}

func (e *EsClient) Ping() int {
	if e.Client != nil {
		ctx := context.Background()
		var url = util.EnvOrElse("ELASTICSEARCH_URL", "http://127.0.0.1:9200")
		info, code, _ := e.Client.Ping(url).Do(ctx)
		if info != nil {
			log.Printf("Elasticsearch returned with code %d and version %s", code, info.Version.Number)
		}

		return code
	} else {
		log.Printf("Elasticsearch error when trying to ping")
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
