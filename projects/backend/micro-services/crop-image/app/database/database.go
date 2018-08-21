package database

import (
	"log"
	"sync"

	mgo "gopkg.in/mgo.v2"
)

type Connection struct {
	DB *mgo.Database
}

var instance *Connection
var once sync.Once

func GetConnection(server string, database string) *Connection {
	once.Do(func() {
		session, err := mgo.Dial(server)
		if err != nil {
			log.Fatal(err)
		}
		instance = &Connection{
			session.DB(database),
		}
	})
	return instance
}
