package mongo

import (
	"github.com/rodrigodealer/whiplash/models"
)

func HealthcheckMongoDb(services []models.HealthcheckServices,
	healthcheck models.HealthcheckStatus,
	connection MongoDb) models.HealthcheckStatus {
	connection.Connect()
	err := connection.Ping()
	var code = 200
	var status = models.Working
	if err != nil {
		code = 500
		status = models.Failed
	}

	services = append(services, models.HealthcheckServices{Name: "mongodb", State: status, Code: code})
	healthcheck.Services = services
	return healthcheck
}
