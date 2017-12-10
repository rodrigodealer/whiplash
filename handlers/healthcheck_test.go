package handlers

import (
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

	handler := http.HandlerFunc(HealthcheckHandler(esClient, mongoClient))
	handler.ServeHTTP(res, req)

	assert.Equal(t, res.Code, 200)
	assert.Equal(t, res.Body.String(), "WORKING")
}

func TestFailedHealthcheck(t *testing.T) {
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	res := httptest.NewRecorder()
	esClient := new(esClientMock)
	mongoClient := new(mongoClientMock)
	esClient.On("Ping").Return(500)

	handler := http.HandlerFunc(HealthcheckHandler(esClient, mongoClient))
	handler.ServeHTTP(res, req)

	assert.Equal(t, res.Code, 500)
	assert.Equal(t, res.Body.String(), "{\"status\":\"FAILED\",\"services\":[{\"name\":\"elasticsearch\",\"state\":\"FAILED\",\"code\":500}]}\n")
}
