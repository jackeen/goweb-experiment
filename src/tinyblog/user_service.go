package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type UserService struct {
	DBC *MDBC
	C   *mgo.Collection
	S   *Session
}

func (self *UserService) Save(user *User) *ResMessage {

	t := &TimeData{}
	t.Now()
	user.Id_ = bson.NewObjectId()
	user.CreateTime = t
	err := self.C.Insert(user)

	return getResMessage(err, SAVE_SUCCESS, USER_MODE_CODE)
}

func (self *UserService) GetList(sel *SelectData) []User {

	ul := make([]User, sel.Limit)
	q := self.C.Find(sel.Condition)
	q.All(&ul)
	return ul
}

func (self *UserService) GetOne(sel *SelectData) *User {

	q := self.C.Find(sel.Condition)
	user := &User{}
	q.One(user)
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

/*type Author struct {
	userGroupList []UserGroup
}

func (self *Author) Init() {

}

func (self *Author) getUsr(req *REQ, s *Session) (bool, *User) {

	var usr *User
	uuid := req.GetOneCookieValue("uuid")

	if uuid == "" {
		return false, usr
	}

	return s.GetCurUsr(uuid)
}

func (self *Author) EditPost(req *REQ, s *Session) bool {

}

func (self *Author) DelPost(req *REQ, s *Session) bool {

}*/
