package helper

import (
	"gopkg.in/mgo.v2"
)

var (
	MongoSession    *mgo.Session
	UsersCollection *mgo.Collection
)
