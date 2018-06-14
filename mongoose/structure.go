package main

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Callback struct {
	Data  interface{} `json:"data"`
	Error error       `json:"error"`
}

type Insert struct {
	Collection *mgo.Collection `json:"collection"`
	Data       interface{}     `json:"data"`
}

type BulkInsert struct {
	Collection *mgo.Collection `json:"collection"`
	Data       []interface{}   `json:"data"`
}

type InsertAsync struct {
	Collection *mgo.Collection `json:"collection"`
	Data       interface{}     `json:"data"`
	Callback   chan *Callback  `json:"callback"`
}

type Update struct {
	Collection *mgo.Collection `json:"collection"`
	Id         bson.ObjectId   `json:"id"`
	Data       interface{}     `json:"data"`
}

type UpdateAll struct {
	Collection *mgo.Collection `json:"collection"`
	Query      bson.M          `json:"query"`
	Data       interface{}     `json:"data"`
}

type FindByID struct {
	Collection *mgo.Collection `json:"collection"`
	Id         bson.ObjectId   `json:"id"`
}

type Find struct {
	Collection *mgo.Collection `json:"collection"`
	Query      bson.M          `json:"query"`
	Options    map[string]int  `json:"options"`
}

type FindAll struct {
	Collection *mgo.Collection `json:"collection"`
}

type Remove struct {
	Collection *mgo.Collection `json:"collection"`
	Query      bson.M          `json:"query"`
}

type RemoveAll struct {
	Collection *mgo.Collection `json:"collection"`
}
