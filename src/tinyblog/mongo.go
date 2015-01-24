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

type BSONM bson.M

type SelectData struct {
	Tab   string
	Sel   BSONM
	Sort  string
	Limit int
	Res   interface{}
	err   error
}

type MDBC struct {
	Host string
	User string
	Pass string
	Name string
	S    *mgo.Session
	DB   *mgo.Database
}

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
		log.Println("mongodb: ", err)
	}

	self.S = s
	self.DB = s.DB(self.Name)
}

func (self *MDBC) Insert(tab string, data interface{}) {

	c := self.DB.C(tab)
	err := c.Insert(data)
	logErr(err)
}

func (self *MDBC) SelectOne(tab string, sel BSONM, res interface{}) {

	c := self.DB.C(tab)
	err := c.Find(sel).One(res)
	logErr(err)
}

func (self *MDBC) SelectArray() {

}

func (self *MDBC) Select(sel *SelectData) {

	c := self.DB.C(sel.Tab)
	query := c.Find(sel.Sel)
	query = query.Sort(sel.Sort)
	query = query.Limit(sel.Limit)
	err := query.All(sel.Res)
	sel.err = err
	logErr(err)
}

func (self *MDBC) UpdateSet(tab string, sel BSONM, data interface{}) {

	c := self.DB.C(tab)
	err := c.Update(sel, bson.M{"$set": data})
	logErr(err)
}

func (self *MDBC) UpdateInc(tab string, sel BSONM, name string, inc int) {

	c := self.DB.C(tab)
	err := c.Update(sel, bson.M{"$inc": bson.M{name: inc}})
	logErr(err)
}

func (self *MDBC) UpdatePush(tab string, sel BSONM, name string, data interface{}) {

	c := self.DB.C(tab)
	err := c.Update(sel, bson.M{"$push": bson.M{name: data}})
	logErr(err)
}

func (self *MDBC) UpdatePull(tab string, sel BSONM, colName string, colSel BSONM) {

	c := self.DB.C(tab)
	err := c.Update(sel, bson.M{"$pull": bson.M{colName: colSel}})
	logErr(err)
}

func (self *MDBC) Delete(tab string, sel BSONM) {

	c := self.DB.C(tab)
	err := c.Remove(sel)
	logErr(err)
}
