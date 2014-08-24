package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

func logErr(err error) {
	if err != nil {
		log.Println("mongo", err)
	}
}

type MDBC struct {
	Host string
	User string
	Pass string
	Name string
	S    *mgo.Session
	DB   *mgo.Database
}

type Selector bson.M

func (self *MDBC) GetDBRef(tab string, _id bson.ObjectId) mgo.DBRef {
	ref := mgo.DBRef{
		Collection: tab,
		Id:         _id,
	}
	return ref
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

func (self *MDBC) Insert(tab string, data interface{}) {

	c := self.DB.C(tab)
	err := c.Insert(data)
	logErr(err)
}

func (self *MDBC) SelectOne(tab string, sel Selector, res interface{}) {

	c := self.DB.C(tab)
	err := c.Find(sel).One(res)
	logErr(err)
}

func (self *MDBC) SelectArray() {

}

func (self *MDBC) Select(tab string, sel Selector, sort string, offset int, limit int, res interface{}) {

	c := self.DB.C(tab)
	query := c.Find(sel)

	if sort != "" {
		query = query.Sort(sort)
	}

	query = query.Skip(offset).Limit(limit)

	err := query.All(res)
	logErr(err)
}

func (self *MDBC) UpdateSet(tab string, sel Selector, data interface{}) {

	c := self.DB.C(tab)
	err := c.Update(sel, bson.M{"$set": data})
	logErr(err)
}

func (self *MDBC) UpdateInc(tab string, sel Selector, name string, inc int) {

	c := self.DB.C(tab)
	err := c.Update(sel, bson.M{"$inc": bson.M{name: inc}})
	logErr(err)
}

func (self *MDBC) UpdatePush(tab string, sel Selector, colName string, data interface{}) {

	c := self.DB.C(tab)
	err := c.Update(sel, bson.M{"$push": bson.M{colName: data}})
	logErr(err)
}

func (self *MDBC) UpdatePull(tab string, sel Selector, colName string, colSel Selector) {

	c := self.DB.C(tab)
	err := c.Update(sel, bson.M{"$pull": bson.M{colName: colSel}})
	logErr(err)
}

func (self *MDBC) Delete(tab string, sel Selector) {

	c := self.DB.C(tab)
	err := c.Remove(sel)
	logErr(err)
}
