package main

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Callback struct {
	Data  interface{}
	Error error
}

type Insert struct {
	Collection *mgo.Collection
	Data       interface{}
}

type BulkInsert struct {
	Collection *mgo.Collection
	Data       []interface{}
}

type InsertAsync struct {
	Collection *mgo.Collection
	Data       interface{}
	Callback   chan *Callback
}

type Update struct {
	Collection *mgo.Collection
	Id         bson.ObjectId
	Data       interface{}
}

type UpdateAll struct {
	Collection *mgo.Collection
	Query      bson.M
	Data       interface{}
}

type FindByID struct {
	Collection *mgo.Collection
	Id         bson.ObjectId
}

type Find struct {
	Collection *mgo.Collection
	Query      bson.M
	Options    map[string]int
}

type FindAll struct {
	Collection *mgo.Collection
}

type Remove struct {
	Collection *mgo.Collection
	Query      bson.M
}

type RemoveAll struct {
	Collection *mgo.Collection
}
