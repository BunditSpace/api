package helper

import (
	"gopkg.in/mgo.v2"
)

var (
	Mongo_session    *mgo.Session
	Users_collection *mgo.Collection
)
