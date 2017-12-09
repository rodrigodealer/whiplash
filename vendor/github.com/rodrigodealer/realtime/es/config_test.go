package es

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigReturnIfNotEmpty(t *testing.T) {
	var value = config("bla")

	assert.Equal(t, value, "bla")
}

func TestConfigReturnIfEmpty(t *testing.T) {
	var value = config("")

	assert.Equal(t, value, "http://127.0.0.1:9200")
}
