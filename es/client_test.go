package es

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFailedConnection(t *testing.T) {
	conn := &EsClient{}
	var oldenv = os.Getenv("ELASTICSEARCH_URL")
	os.Setenv("ELASTICSEARCH_URL", "localhost:9400")
	assert.Panics(t, func() { conn.Connect() })
	os.Setenv("ELASTICSEARCH_URL", oldenv)
}
