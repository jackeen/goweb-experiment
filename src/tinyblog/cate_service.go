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

	if c.Name == "" {
		return getUserResMessage(false, REQUIRED_DEFAULT, CATE_MODE_CODE)
	}
	err := self.C.Insert(c)
	return getResMessage(err, SAVE_SUCCESS, CATE_MODE_CODE)
}

func (self *CateService) GetOne(name string) *Cate {
	c := new(Cate)
	self.C.Find(bson.M{"name": name}).One(c)
	return c
}

func (self *CateService) IsExist(name string) bool {
	c := self.GetOne(name)
	if c.Name == "" {
		return false
	} else {
		return true
	}
}

func (self *CateService) GetList(pName string) []Cate {
	q := self.C.Find(bson.M{"parent": pName})
	n, err := q.Count()
	if err != nil {
		n = 0
	}
	cateList := make([]Cate, n)
	q.All(&cateList)
	return cateList
}

func (self *CateService) GetNames(pName string) []string {
	cList := self.GetList(pName)
	l := len(cList)
	nameList := make([]string, l)
	for i := 0; i < l; i++ {
		nameList[i] = cList[i].Name
	}
	return nameList
}

func (self *CateService) Update(name string, c *Cate) *ResMessage {
	err := self.C.Update(bson.M{"name": name}, bson.M{"$set": c})
	return getResMessage(err, UPDATE_SUCCESS, CATE_MODE_CODE)
}

func (self *CateService) Del(name string) *ResMessage {

	c := self.GetOne(name)
	children := c.Children
	l := len(children)
	for i := 0; i < l; i++ {
		cName := children[i]
		c := self.GetOne(cName)
		c.Parent = ""
		self.Update(c.Name, c)
	}
	err := self.C.Remove(bson.M{"name": name})
	return getResMessage(err, DEL_SUCCESS, CATE_MODE_CODE)
}
