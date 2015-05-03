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

func (self *UserService) GetList(sel *SelectData) *UserList {
	ul := &UserList{}
	q := self.C.Find(sel.Condition)
	err := q.All(ul)
	findPanic(err)
	return ul
}

func (self *UserService) GetOne(sel *SelectData) *User {

	q := self.C.Find(sel.Condition)
	user := &User{}
	q.One(user)
	//findPanic(err)
	return user
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


*/
