package models

import (
	db "../helper/db"
	"gopkg.in/mgo.v2/bson"
)

//User object user data
type User struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Firstname string        `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string        `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Username  string        `json:"username,omitempty" bson:"username,omitempty"`
	Password  string        `json:"password,omitempty" bson:"password,omitempty"`
	Image     string        `json:"image,omitempty" bson:"image,omitempty"`
	Detail    string        `json:"detail,omitempty" bson:"detail,omitempty"`
}

//SaveToDB s
func (u *User) SaveToDB() error {
	err := db.UsersCollection.Insert(&u)

	if err != nil {
		return err
	}
	return nil
}

//ReadFromDB r
func (u *User) ReadFromDB() ([]User, error) {
	result := []User{}
	err := db.UsersCollection.Find(nil).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//ReadByID only one
func (u *User) ReadByID() (*User, error) {
	err := db.UsersCollection.Find(bson.M{"_id": u.ID}).One(&u)
	if err != nil {
		return nil, err
	}
	return u, nil
}