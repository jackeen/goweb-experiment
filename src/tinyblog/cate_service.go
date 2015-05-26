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

func (self *CateService) GetRootList() []Cate {

	var cateList []Cate
	q := self.C.Find(bson.M{"parent": ""})
	n, err := q.Count()
	if err == nil {
		cateList = make([]Cate, n)
	} else {
		cateList = make([]Cate, 0)
	}
	return cateList
}

func (self *CateService) GetOne(n string) *Cate {
	c := new(Cate)
	self.C.Find(bson.M{"name": n}).One(c)
	return c
}

func (self *CateService) ChangeName(c *Cate, n string) *ResMessage {
	c.Name = n
	err := self.C.Update(bson.M{"name": c.Name}, c)
	return getResMessage(err, DEL_SUCCESS, CATE_MODE_CODE)
}

func (self *CateService) ChangeParent(c *Cate) *ResMessage {
	err := self.C.Update(bson.M{"name": c.Name}, bson.M{"$set": c})
	return getResMessage(err, UPDATE_SUCCESS, CATE_MODE_CODE)
}

func (self *CateService) Del(n string) *ResMessage {

	c := self.GetOne(n)
	children := c.Children
	l := len(children)
	for i := 0; i < l; i++ {
		name := children[i]
		c := self.GetOne(name)
		c.Parent = ""
		self.ChangeParent(c)
	}
	err := self.C.Remove(bson.M{"name": n})
	return getResMessage(err, DEL_SUCCESS, CATE_MODE_CODE)
}
