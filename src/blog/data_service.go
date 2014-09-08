package main

import (
	//"log"
	//"reflect"
	//"errors"
	//"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

const (
	NUM_TAB    = "num"
	POST_TAB   = "post"
	CATE_TAB   = "cate"
	USER_TAB   = "user"
	TAG_TAB    = "tag"
	NAV_TAB    = "nav"
	CONFIG_TAB = "config"
)

//inc id num data I/O
type NumService struct{}

func (self *NumService) Init(dbc *MDBC) {
	dbc.Insert(NUM_TAB, &Num{0, 0, 0, 0})
}

func (self *NumService) incId(dbc *MDBC, colName string, i int) *Num {
	res := &Num{}
	dbc.UpdateInc(NUM_TAB, nil, colName, i)
	dbc.SelectOne(NUM_TAB, nil, res)
	return res
}

var IncNum *NumService = new(NumService)

//user data IO
type UserService struct{}

func (self *UserService) Insert(dbc *MDBC, name string, pass string, nick string, email string) {

	user := &User{
		Id_:        bson.NewObjectId(),
		Id:         IncNum.incId(dbc, "user", 1).User,
		Name:       name,
		Pass:       pass,
		Nick:       nick,
		Email:      email,
		CreateTime: time.Now(),
	}
	dbc.Insert(USER_TAB, user)
}

func (self *UserService) Update(dbc *MDBC, sel Selector, data interface{}) {
	dbc.UpdateSet(USER_TAB, sel, data)
}

func (self *UserService) Select(dbc *MDBC) {

}

func (self *UserService) Delete(dbc *MDBC) {

}

func (self *UserService) HasUser(dbc *MDBC, name string) bool {
	res := &User{}
	dbc.SelectOne(USER_TAB, Selector{"name": name}, res)
	if res.Name == "" {
		return false
	} else {
		return true
	}
}

//
type CateService struct{}

func (self *CateService) Inert(dbc *MDBC) {

}
