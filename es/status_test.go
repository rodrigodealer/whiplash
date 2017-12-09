package es

import (
	"testing"

	"github.com/rodrigodealer/whiplash/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSuccessfulHealthcheck(t *testing.T) {
	var services []models.HealthcheckServices
	var healthcheck = models.HealthcheckStatus{Status: Working, Services: services}
	client := new(clientMock)
	client.On("Ping").Return(200)

	var status = HealthcheckElasticsearch(services, healthcheck, client)

	assert.Equal(t, status.Status, Working)

}

func TestFailedHealthcheck(t *testing.T) {
	var services []models.HealthcheckServices
	var healthcheck = models.HealthcheckStatus{Status: Working, Services: services}
	client := new(clientMock)
	client.On("Ping").Return(500)

	var status = HealthcheckElasticsearch(services, healthcheck, client)

	assert.Equal(t, status.Status, Failed)
}

type clientMock struct {
	mock.Mock
}

func (o clientMock) Ping() int {
	args := o.Called()
	return args.Int(0)
}

func (o clientMock) Connect() {
}
