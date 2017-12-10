package es

import (
	"github.com/rodrigodealer/whiplash/models"
)

func HealthcheckElasticsearch(services []models.HealthcheckServices,
	healthcheck models.HealthcheckStatus,
	connection ElasticSearch) models.HealthcheckStatus {
	var code = connection.Ping()
	var status = models.Working
	if code != 200 {
		status = models.Failed
	}
	services = append(services, models.HealthcheckServices{Name: "elasticsearch", State: status, Code: code})
	healthcheck.Status = status
	healthcheck.Services = services
	return healthcheck
}
