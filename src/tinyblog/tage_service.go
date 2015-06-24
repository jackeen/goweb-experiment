package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type TageService struct {
	DBC *MDBC
	C   *mgo.Collection
	S   *Session
}

func (self *TageService) Save(t *Tag) *ResMessage {
	err := self.C.Insert(t)
	return getResMessage(err, SAVE_SUCCESS, TAGE_MODE_CODE)
}

func (self *TageService) GetList() []Tag {
	q := self.C.Find(bson{})
	n, err := q.Count()
	if err != nil {
		n = 0
	}
	tagList := make([]Tag, n)
	q.All(&tagList)
	return tagList
}

func (self *TageService) Del(name string) *ResMessage {
	err := self.C.Remove(bson{"name": name})
	return getResMessage(err, DEL_SUCCESS, TAGE_MODE_CODE)
}
