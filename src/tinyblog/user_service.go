package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type UserService struct {
	DBC *MDBC
	C   *mgo.Collection
	S   *Session
}

func (self *UserService) Save(user *User) *ResMessage {

	user.Id_ = bson.NewObjectId()
	user.CreateTime = time.Now()
	err := self.C.Insert(user)

	return getResMessage(err, SAVE_SUCCESS, USER_MODE_CODE)
}

func (self *UserService) GetList(sel *SelectData) {

}

func (self *UserService) Login(sel *SelectData) {

	n := sel.Condition["name"]
	p := sel.Condition["pass"]

	q := self.C.Find(bson.M{"name": n, "pass": p})
	err := q.One(sel.Err)
}

/*
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
*/
