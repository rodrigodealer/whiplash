package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodrigodealer/whiplash/es"
	"github.com/rodrigodealer/whiplash/handlers"
)

func Server() http.Handler {
	conn := &es.EsClient{}
	conn.Connect()

	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", handlers.HealthcheckHandler(conn)).Name("/healthcheck").Methods("GET")
	return r
}
