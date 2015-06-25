package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type TagService struct {
	DBC *MDBC
	C   *mgo.Collection
	S   *Session
}

func (self *TagService) Save(t *Tag) *ResMessage {
	err := self.C.Insert(t)
	return getResMessage(err, SAVE_SUCCESS, TAGE_MODE_CODE)
}

func (self *TagService) GetList() []Tag {
	q := self.C.Find(bson.M{})
	n, err := q.Count()
	if err != nil {
		n = 0
	}
	tagList := make([]Tag, n)
	q.All(&tagList)
	return tagList
}

func (self *TagService) GetOne(name string) *Tag {
	t := new(Tag)
	self.C.Find(bson.M{"name": name}).One(t)
	return t
}

func (self *TagService) IsExist(name string) bool {
	t := self.GetOne(name)
	if t.Name == "" {
		return false
	} else {
		return true
	}
}

func (self *TagService) Del(name string) *ResMessage {
	err := self.C.Remove(bson.M{"name": name})
	return getResMessage(err, DEL_SUCCESS, TAGE_MODE_CODE)
}
