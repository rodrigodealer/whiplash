package handlers

import (
	"fmt"
	"net/http"

	"github.com/rodrigodealer/realtime/es"
	"github.com/rodrigodealer/realtime/models"
	"github.com/rodrigodealer/realtime/services"
	"github.com/rodrigodealer/realtime/tracing"
)

func FacebookUpdateHandler(connection es.ElasticSearch,
	fbclient services.FacebookClient) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			parent := tracing.Trace("Facebook API")
			var facebookUpdate = models.FacebookUpdate{}
			facebookUpdate.FromRequest(r.Body)
			token := services.GetToken(fbclient, parent)
			for _, entry := range facebookUpdate.Entry {
				getUserAndIndex(entry, token, fbclient, connection, parent)
			}
			defer parent.Finish()
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello, %v", r.URL.Query().Get("u"))
	}
}

func getUserAndIndex(entry models.FacebookUpdateEntry,
	token *models.FacebookToken, client services.FacebookClient,
	conn es.ElasticSearch, parent *tracing.Tracing) {
	var response = services.GetUid(entry, token, client, parent)
	facebookUser := &models.FacebookUser{}
	facebookUser.FromJson(response.Body)
	go conn.IndexUser("facebook", facebookUser, parent)
}
