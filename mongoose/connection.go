package main

import (
	"gopkg.in/mgo.v2"
)

type Config struct {
	Connection  string //connection url i.e, localhost:27017
	Database    string //database name
	DialInfo    *mgo.DialInfo
	Session     *mgo.Session
	Collections map[string]*mgo.Collection
}

var c *mgo.Collection = nil

func (config *Config) Connect() (*Config, error) {

	session, err := mgo.Dial(config.Connection)

	if err == nil {
		config.Session = session
		config.Collections = make(map[string]*mgo.Collection)
		config.Session.DB(config.Database)
	}
	return config, err
}

func Close(session *mgo.Session) {
	session.Close()
}
