package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type CateService struct {
	DBC *MDBC
	C   *mgo.Collection
	S   *Session
}

func (self *CateService) Save(c *Cate) *ResMessage {

	if c.Name == "" || c.Parent == "" {
		return getUserResMessage(false, REQUIRED_DEFAULT, CATE_MODE_CODE)
	}
	err := self.C.Insert(c)
	return getResMessage(err, SAVE_SUCCESS, CATE_MODE_CODE)
}

func (self *CateService) GetRoot() []Cate {

	cateList := make([]Cate, 10)
	self.C.Find(bson.M{"parent": ""}).All(cateList)
	return cateList
}

func (self *CateService) FindOne(name string) *Cate {
	c := new(Cate)
	self.C.Find(bson.M{"name": name}).One(c)
	return c
}

//func (self *CateService) Delete(name string) *ResMessage {

//err := self.C.Remove(bson.M{"name": name})
//}
