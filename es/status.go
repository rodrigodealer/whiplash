package es

import (
	"github.com/rodrigodealer/whiplash/models"
)

const (
	Failed  = "FAILED"
	Working = "WORKING"
)

func HealthcheckElasticsearch(services []models.HealthcheckServices,
	healthcheck models.HealthcheckStatus,
	connection ElasticSearch) models.HealthcheckStatus {
	var code = connection.Ping()
	if code != 200 {
		services = append(services, models.HealthcheckServices{Name: "elasticsearch", State: Failed, Code: code})
		healthcheck.Status = Failed
		healthcheck.Services = services
	}
	return healthcheck
}
