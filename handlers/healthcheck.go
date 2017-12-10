package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rodrigodealer/whiplash/es"
	"github.com/rodrigodealer/whiplash/models"
	"github.com/rodrigodealer/whiplash/mongo"
)

func HealthcheckHandler(connection es.ElasticSearch,
	mongoConnection mongo.MongoDb) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var services []models.HealthcheckServices
		var healthcheck = models.HealthcheckStatus{Status: models.Working, Services: services}
		healthcheck = es.HealthcheckElasticsearch(healthcheck.Services, healthcheck, connection)
		healthcheck = mongo.HealthcheckMongoDb(healthcheck.Services, healthcheck, mongoConnection)
		if healthcheck.Status == models.Failed {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(healthcheck)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(healthcheck)
		}
	}
}
