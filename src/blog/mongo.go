package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

type BSON bson.M
type MongoId bson.ObjectId

type MDBC struct {
	Host string
	User string
	Pass string
	Name string
	S    *mgo.Session
	DB   *mgo.Database
}

/*
func (self *MDBC) GetMongoId() MongoId {
	return bson.NewObjectId()
}*/

func (self *MDBC) Init() {

	dbQuery := "mongodb://" + self.User + ":" + self.Pass + "@" + self.Host + "/" + self.Name
	s, err := mgo.Dial(dbQuery)
	if err != nil {
		panic(err)
	}

	self.S = s
	self.DB = s.DB(self.Name)
}

func (self *MDBC) Insert(tab string, data interface{}) {

	c := self.DB.C(tab)
	err := c.Insert(data)
	if err != nil {
		log.Println(err)
	}
}

func (self *MDBC) Select(tab string, selector BSON, sort string, offset int, limit int, res interface{}) {

	c := self.DB.C(tab)

	query := c.Find(selector)
	query = query.Sort(sort).Skip(offset).Limit(limit)
	query.All(res)
}

func (self *MDBC) UpdateSet(tab string, selector BSON, data interface{}) {

	c := self.DB.C(tab)
	err := c.Update(selector, BSON{"$set": data})
	if err != nil {
		panic(err)
	}
}

func (self *MDBC) UpdateInc(tab string, selector BSON, name string, inc int) {

	c := self.DB.C(tab)
	err := c.Update(selector, BSON{"$inc": BSON{name: inc}})
	if err != nil {
		panic(err)
	}
}
