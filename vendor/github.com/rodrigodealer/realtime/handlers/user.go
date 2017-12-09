package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rodrigodealer/realtime/es"
)

func UserHandler(connection es.ElasticSearch) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := connection.GetUser("facebook", r.URL.Query().Get("u"))

		w.Header().Set("Content-Type", "application/json")
		if err == nil {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
		json.NewEncoder(w).Encode(user)
	}
}
