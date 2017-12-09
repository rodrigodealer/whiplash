package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFacebookUser(t *testing.T) {
	fbuser := FacebookUser{ID: "Take Five", Name: "Bla"}

	assert.Equal(t, fbuser.ID, "Take Five")
	assert.Equal(t, fbuser.Name, "Bla")
}

func TestFacebookUserToJson(t *testing.T) {
	fbuser := FacebookUser{ID: "Take Five", Name: "Bla"}

	assert.Equal(t, fbuser.ToJson(), "{\"id\":\"Take Five\",\"name\":\"Bla\"}")
}
func TestFacebookUserUnmarshall(t *testing.T) {
	jsonText := "{\"id\":\"user\",\"name\":\"user\"}"
	facebookUser := &FacebookUser{}
	facebookUser.FromJson(jsonText)
	assert.Equal(t, "user", facebookUser.ID)
	assert.Equal(t, "user", facebookUser.Name)
}
