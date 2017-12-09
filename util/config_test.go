package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigReturnIfNotEmpty(t *testing.T) {
	var value = Config("bla")

	assert.Equal(t, value, "bla")
}

func TestConfigReturnIfEmpty(t *testing.T) {
	var value = Config("")

	assert.Equal(t, value, "http://127.0.0.1:9200")
}

func TestEnvironmentVariableNotPresent(t *testing.T) {
	var value = "myenv"
	var defaultValue = "myval"

	var actualValue = EnvOrElse(value, defaultValue)
	assert.Equal(t, actualValue, defaultValue)
}

func TestEnvironmentVariableIsPresent(t *testing.T) {
	var value = "myenv"
	var otherValue = "othervalue"
	var defaultValue = "myval"

	os.Setenv(value, "othervalue")

	var actualValue = EnvOrElse(value, defaultValue)
	assert.Equal(t, actualValue, otherValue)

	os.Setenv(value, "")
}
