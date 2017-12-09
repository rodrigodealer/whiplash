package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rodrigodealer/realtime/models"
	"github.com/stretchr/testify/assert"
)

func TestUserGet(t *testing.T) {
	req, _ := http.NewRequest("GET", "/user?u=chris", nil)
	res := httptest.NewRecorder()
	client := new(clientMock)
	user := models.FacebookUser{Name: "User", ID: "1"}
	client.On("GetUser").Return(user, nil)
	handler := http.HandlerFunc(UserHandler(client))

	handler.ServeHTTP(res, req)

	assert.Equal(t, res.Body.String(), "{\"id\":\"\",\"name\":\"\"}\n")
}

func TestUserGetWithNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", "/user?u=chris", nil)
	res := httptest.NewRecorder()
	client := new(clientMock)
	user := models.FacebookUser{Name: "User", ID: "1"}
	client.On("GetUser").Return(user, errors.New("Error"))
	handler := http.HandlerFunc(UserHandler(client))

	handler.ServeHTTP(res, req)

	assert.Equal(t, res.Body.String(), "{\"id\":\"\",\"name\":\"\"}\n")
}
