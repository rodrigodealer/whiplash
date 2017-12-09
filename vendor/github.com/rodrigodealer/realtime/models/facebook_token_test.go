package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFacebookTokenUnmarshall(t *testing.T) {
	jsonText := "{\"access_token\":\"1203725573077981|HrXuRe8Ph7hzzGHH-8ATjMMlojg\",\"token_type\":\"bearer\"}"
	facebookToken := &FacebookToken{}
	facebookToken.FromJson(jsonText)
	assert.Equal(t, "bearer", facebookToken.Type)
	assert.Equal(t, "1203725573077981|HrXuRe8Ph7hzzGHH-8ATjMMlojg", facebookToken.Token)
}
