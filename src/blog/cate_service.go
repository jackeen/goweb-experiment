package main

import (
	//"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type CateService struct {
	NumService
	DBC *MDBC
}

func (self *CateService) Inert(name string, exp string, pid int) {

	cate := &Cate{
		Id_:      bson.NewObjectId(),
		Id:       self.incId(self.DBC, "cate", 1).Cate,
		Name:     name,
		Explain:  exp,
		ParentId: pid,
	}
	self.DBC.Insert(CATE_TAB, cate)
}

func (self *CateService) Select(id int, cate *Cate) {

	self.DBC.SelectOne(CATE_TAB, BSONM{"id": id}, cate)
}

func (self *CateService) Update(id int, data interface{}) {

	self.DBC.UpdateSet(CATE_TAB, BSONM{"id": id}, data)
}

func (self *CateService) Delete(id int) {
	self.DBC.Delete(CATE_TAB, BSONM{"id": id})
}
