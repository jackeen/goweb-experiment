package main

import (
//"labix.org/v2/mgo"
//"labix.org/v2/mgo/bson"
//"time"
)

const (
	POST_TAB   = "post"
	CATE_TAB   = "cate"
	USER_TAB   = "user"
	TAG_TAB    = "tag"
	CONFIG_TAB = "config"
)

const (
	SaveSuccess  = "save success"
	SaveFail     = "save fail"
	SaveDataFail = "save data fail"
)

type ResMessage struct {
	State   bool
	Message string
}

type ResData struct {
	State bool
	Count int
	Data  interface{}
}

type DataService struct {
	DBC  *MDBC
	Post *PostService
	User *UserService
}

func (self *DataService) Init(dbc *MDBC) {
	self.DBC = dbc
	self.Post = &PostService{
		DBC: dbc,
		C:   dbc.DB(POST_TAB),
	}
	self.User = &UserService{
		DBC: dbc,
		C:   dbc.DB(USER_TAB),
	}
}
