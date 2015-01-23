package main

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type UserService struct {
	NumService
	DBC *MDBC
}

func (self *UserService) Insert(name string, pass string, nick string, email string, power int) {

	user := &User{
		Id_:        bson.NewObjectId(),
		Id:         self.incId(self.DBC, "user", 1).User,
		Name:       name,
		Pass:       pass,
		Nick:       nick,
		Email:      email,
		Power:      power,
		CreateTime: time.Now(),
	}
	self.DBC.Insert(USER_TAB, user)
}

func (self *UserService) Update(sel BSONM, data interface{}) {
	self.DBC.UpdateSet(USER_TAB, sel, data)
}

func (self *UserService) Select(dbc *MDBC) {

}

func (self *UserService) Delete(dbc *MDBC) {

}

func (self *UserService) LoginSelect(u string, p string, res *User) {
	sel := BSONM{
		"name": u,
		"pass": p,
	}
	self.DBC.SelectOne(USER_TAB, sel, res)
}

func (self *UserService) HasUser(name string) bool {
	res := &User{}
	self.DBC.SelectOne(USER_TAB, BSONM{"name": name}, res)
	if res.Name == "" {
		return false
	} else {
		return true
	}
}
