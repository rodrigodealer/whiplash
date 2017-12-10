package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessfulHealthcheck(t *testing.T) {
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	res := httptest.NewRecorder()
	esClient := new(esClientMock)
	mongoClient := new(mongoClientMock)
	esClient.On("Ping").Return(200)
	mongoClient.On("Ping").Return(nil)

	handler := http.HandlerFunc(HealthcheckHandler(esClient, mongoClient))
	handler.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "{\"status\":\"WORKING\",\"services\":[{\"name\":\"elasticsearch\",\"state\":\"WORKING\",\"code\":200},{\"name\":\"mongodb\",\"state\":\"WORKING\",\"code\":200}]}\n", res.Body.String())
}

func TestFailedHealthcheck(t *testing.T) {
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	res := httptest.NewRecorder()
	esClient := new(esClientMock)
	mongoClient := new(mongoClientMock)
	esClient.On("Ping").Return(500)
	mongoClient.On("Ping").Return(errors.New("Session invalid"))

	handler := http.HandlerFunc(HealthcheckHandler(esClient, mongoClient))
	handler.ServeHTTP(res, req)

	assert.Equal(t, 500, res.Code)
	assert.Equal(t, "{\"status\":\"FAILED\",\"services\":[{\"name\":\"elasticsearch\",\"state\":\"FAILED\",\"code\":500},{\"name\":\"mongodb\",\"state\":\"FAILED\",\"code\":500}]}\n", res.Body.String())
}
