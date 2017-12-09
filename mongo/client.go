package mongo

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

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

func (m *MongoConnection) Close() {
	m.Session.Close()
}
