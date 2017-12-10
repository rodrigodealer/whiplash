package mongo

import (
	"errors"
	"log"

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
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Panic("Could not connect to MongoDB:\n", err.Error())
	}
	session.SetMode(mgo.Monotonic, true)
	m.Session = session
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
