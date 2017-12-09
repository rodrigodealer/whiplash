package services

import (
	"bytes"
	"errors"
	"testing"

	"github.com/rodrigodealer/realtime/models"
	"github.com/rodrigodealer/realtime/tracing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacebookTokenUrl(t *testing.T) {
	url := FacebookTokenUrl("123", "123")
	assert.Equal(t, url, "https://graph.facebook.com/oauth/access_token?client_id=123&client_secret=123&grant_type=client_credentials")
}

func TestFacebookUidUrl(t *testing.T) {
	url := FacebookUIDUrl("123", "123")
	assert.Equal(t, "https://graph.facebook.com/v2.10/123?fields=id,name,picture{url}&access_token=123", url)
}

func TestGetUid(t *testing.T) {
	fbClient := new(fbClientMock)
	fbClient.On("GetRequest").
		Return("{\"id\":\"1449629478400719\",\"name\":\"Rodrigo Oliveira\"}", 200)
	var list = models.FacebookUpdateEntry{UID: "123"}
	var token = &models.FacebookToken{Token: "123"}

	var result = GetUid(list, token, fbClient, &tracing.Tracing{})

	assert.Equal(t, 200, result.Code)
}

func TestGetToken(t *testing.T) {
	fbClient := new(fbClientMock)
	fbClient.On("GetRequest").Return("{\"access_token\":\"1449629478400719\",\"token_type\":\"bearer\"}", 200)

	var result = GetToken(fbClient, &tracing.Tracing{})
	assert.Equal(t, "1449629478400719", result.Token)
	assert.Equal(t, "bearer", result.Type)
}

func TestHttpResponseBodyToString(t *testing.T) {
	body := []byte("mybody")
	reader := bytes.NewReader(body)

	newbody := HttpResponseBodyToString(reader)
	assert.Equal(t, "mybody", newbody)
}

type fbClientMock struct {
	mock.Mock
}

func (o fbClientMock) GetRequest(url string) (FbResponse, error) {
	args := o.Called()
	var response = FbResponse{Code: args.Int(1), Body: args.String(0)}
	return response, errors.New("can't work with 42")
}
