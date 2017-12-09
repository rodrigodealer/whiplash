package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodrigodealer/whiplash/es"
	"github.com/rodrigodealer/whiplash/handlers"
	"github.com/rodrigodealer/whiplash/mongo"
)

func Server() http.Handler {
	conn := &es.EsClient{}
	conn.Connect()

	mongo := mongo.MongoConnection{}
	mongo.Connect()
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", handlers.HealthcheckHandler(conn, mongo)).Name("/healthcheck").Methods("GET")
	return r
}
