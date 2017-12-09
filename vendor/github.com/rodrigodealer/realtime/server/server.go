package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodrigodealer/realtime/es"
	"github.com/rodrigodealer/realtime/handlers"
	"github.com/rodrigodealer/realtime/services"
)

func Server() http.Handler {
	conn := &es.EsClient{}
	conn.Connect()
	conn.IndexSetup("facebook")

	fbClient := &services.FbClient{}

	r := mux.NewRouter()
	r.HandleFunc("/subscription", handlers.SubscriptionHandler).Name("/subscription").Methods("GET")
	r.HandleFunc("/subscription", handlers.FacebookUpdateHandler(conn, fbClient)).Name("/subscription").Methods("POST")
	r.HandleFunc("/user", handlers.UserHandler(conn)).Name("/user").Methods("GET")
	r.HandleFunc("/healthcheck", handlers.HealthcheckHandler(conn)).Name("/healthcheck").Methods("GET")
	return r
}
