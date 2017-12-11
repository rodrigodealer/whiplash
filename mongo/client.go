package mongo

import (
	"errors"
	"log"
	"time"

	"github.com/rodrigodealer/whiplash/util"
	mgo "gopkg.in/mgo.v2"
)

type MongoDb interface {
	Connect()
	Ping() error
}

type MongoConnection struct {
	Session *mgo.Session
}

func (m *MongoConnection) Connect() {
	session, err := mgo.DialWithTimeout(util.EnvOrElse("MONGO_HOST", "localhost:27017"), 200*time.Millisecond)
	if err != nil {
		log.Print("Could not connect to MongoDB:\n", err.Error())
	}
	if session != nil {
		session.SetMode(mgo.Monotonic, true)
		m.Session = session
	}
}

func (m *MongoConnection) Ping() error {
	if m.Session != nil {
		return m.Session.Ping()
	}
	return errors.New("Session invalid")
}

func (m *MongoConnection) Close() {
	m.Session.Close()
}
