package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rodrigodealer/whiplash/es"
	"github.com/rodrigodealer/whiplash/models"
)

func HealthcheckHandler(connection es.ElasticSearch) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var services []models.HealthcheckServices
		var healthcheck = models.HealthcheckStatus{Status: es.Working, Services: services}
		healthcheck = es.HealthcheckElasticsearch(services, healthcheck, connection)
		w.Header().Set("Content-Type", "application/json")
		if healthcheck.Status == es.Failed {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(healthcheck)
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, es.Working)
		}
	}
}
