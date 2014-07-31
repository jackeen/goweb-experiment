package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Selector bson.M
type MongoId bson.ObjectId

type MDBC struct {
	Host string
	User string
	Pass string
	Name string
	S    *mgo.Session
	DB   *mgo.Database
}

func (self *MDBC) Init() {

	dbQuery := "mongodb://" + self.User + ":" + self.Pass + "@" + self.Host + "/" + self.Name
	s, err := mgo.Dial(dbQuery)
	if err != nil {
		panic(err)
	}

	self.S = s
	self.DB = s.DB(self.Name)
}
