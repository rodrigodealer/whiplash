package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func TestSuccessfulSubscriptionPostWithNilBody(t *testing.T) {
	req, _ := http.NewRequest("POST", "/subscription?u=chris", nil)
	res := httptest.NewRecorder()
	client := new(clientMock)
	fbClient := new(fbClientMock)
	handler := http.HandlerFunc(FacebookUpdateHandler(client, fbClient))

	handler.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, res.Body.String(), "Hello, chris")
}

func TestSuccessfulSubscriptionPostWithBody(t *testing.T) {

	var jsonStr = []byte(`{"object":"user","entry":[{"uid":"100000610422894","id":"100000610422894","time":1232313,"changed_fields":["friends"]}]}`)
	req, _ := http.NewRequest("POST", "/subscription?u=chris", bytes.NewBuffer(jsonStr))
	res := httptest.NewRecorder()
	client := new(clientMock)
	fbClient := new(fbClientMock)
	handler := http.HandlerFunc(FacebookUpdateHandler(client, fbClient))

	handler.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, res.Body.String(), "Hello, chris")
}

func TestSuccessfulSubscriptionPostWithWrongJsonInBody(t *testing.T) {

	var jsonStr = []byte(`{"object": false}`)
	req, _ := http.NewRequest("POST", "/subscription?u=chris", bytes.NewBuffer(jsonStr))
	res := httptest.NewRecorder()
	client := new(clientMock)
	fbClient := new(fbClientMock)
	handler := http.HandlerFunc(FacebookUpdateHandler(client, fbClient))

	handler.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, res.Body.String(), "Hello, chris")
}
