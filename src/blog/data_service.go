package main

import (
	//"log"
	//"reflect"
	//"errors"
	//"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
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

func (self *UserService) Find(dbc *MDBC) {

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

//post data I/O
type PostService struct{}

func (self *PostService) Insert(dbc *MDBC, title string, content string) {

	incId := IncNum.incId(dbc, "post", 1).Post
	currentTime := time.Now()

	data := &Post{
		Id_:          bson.NewObjectId(),
		Id:           incId,
		Title:        title,
		Content:      content,
		Auth:         "admin",
		Cate:         -1,
		Tags:         "",
		CreateTime:   currentTime,
		LastEditTime: currentTime,
	}
	dbc.Insert(POST_TAB, data)
}

func (self *PostService) Find(dbc *MDBC, sel Selector, sort string, offset int, limit int, res *[]Post) {

	dbc.Select(POST_TAB, sel, sort, offset, limit, res)
}

func (self *PostService) Update(dbc *MDBC, sel Selector, data interface{}) {
	dbc.UpdateSet(POST_TAB, sel, data)
}

//
//type
