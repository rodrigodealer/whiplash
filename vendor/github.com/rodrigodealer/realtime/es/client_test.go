package es

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/rodrigodealer/realtime/models"
	"github.com/rodrigodealer/realtime/tracing"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	conn := &EsClient{}
	conn.Connect()

	var ping = conn.Ping()
	assert.Equal(t, 200, ping)
}

func TestFailedPing(t *testing.T) {
	conn := &EsClient{}

	var ping = conn.Ping()
	assert.Equal(t, 500, ping)
}

func TestFailedConnection(t *testing.T) {
	conn := &EsClient{}
	var oldenv = os.Getenv("ELASTICSEARCH_URL")
	os.Setenv("ELASTICSEARCH_URL", "localhost:9400")
	assert.Panics(t, func() { conn.Connect() })
	os.Setenv("ELASTICSEARCH_URL", oldenv)
}

func TestIndexCreation(t *testing.T) {
	ctx := context.Background()
	conn := &EsClient{}
	conn.Connect()
	var now = strconv.Itoa(time.Now().Nanosecond())
	var indexName = "testindex_" + now
	existed, _ := conn.Client.IndexExists(indexName).Do(ctx)
	conn.IndexSetup("testindex_" + now)
	existes, _ := conn.Client.IndexExists(indexName).Do(ctx)

	assert.False(t, existed, "Index shoult not exists")
	assert.True(t, existes, "Index shoult exists")

	conn.Client.DeleteIndex(indexName)
}

func TestIndexRecreation(t *testing.T) {
	ctx := context.Background()
	conn := &EsClient{}
	conn.Connect()
	var now = strconv.Itoa(time.Now().Nanosecond())
	var indexName = "testindex_" + now
	conn.IndexSetup("testindex_" + now)
	existed, _ := conn.Client.IndexExists(indexName).Do(ctx)
	conn.IndexSetup("testindex_" + now)

	assert.True(t, existed, "Index shoult exists already")
	conn.Client.DeleteIndex(indexName)
}

func TestIndexUserAndGetIt(t *testing.T) {
	conn := &EsClient{}
	conn.Connect()
	var now = strconv.Itoa(time.Now().Nanosecond())
	var indexName = "testindex_" + now
	var userID = "1"
	var userName = "Test"
	facebookUser := &models.FacebookUser{ID: userID, Name: userName}
	trac := &tracing.Tracing{}
	conn.IndexUser(indexName, facebookUser, trac)
	time.Sleep(1 * time.Second)

	var retrievedUser, _ = conn.GetUser(indexName, userID)
	assert.Equal(t, userID, retrievedUser.ID)
	assert.Equal(t, userName, retrievedUser.Name)
}

func TestTrytoFindUknownUser(t *testing.T) {
	conn := &EsClient{}
	conn.Connect()
	var now = strconv.Itoa(time.Now().Nanosecond())
	var indexName = "testindex_" + now
	var userID = "2"
	conn.IndexSetup(indexName)

	var _, err = conn.GetUser(indexName, userID)

	assert.Equal(t, "No user found", err.Error())
}
