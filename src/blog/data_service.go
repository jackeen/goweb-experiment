package main

import (
	//"log"
	//"reflect"
	//"errors"
	"labix.org/v2/mgo/bson"
	"time"
)

//increase id function
func increaseNumId(dbc *MDBC, colName string, i int) *Num {

	res := &Num{}
	dbc.UpdateInc(NUM_TAB, nil, colName, i)
	dbc.SelectOne(NUM_TAB, nil, res)
	return res
}

//post data I/O
type PostService struct {
}

func (self *PostService) Insert(dbc *MDBC, title string, content string) {

	var incId int = increaseNumId(dbc, "post", 1).Post

	data := &Post{
		Id_:        bson.NewObjectId(),
		Id:         incId,
		Title:      title,
		Content:    content,
		Auth:       "admin",
		Cate:       -1,
		Tags:       "",
		CreateTime: time.Now(),
	}
	dbc.Insert(POST_TAB, data)
}

func (self *PostService) Find(dbc *MDBC, sel Selector, sort string, offset int, limit int, res *[]Post) {

	dbc.Select(POST_TAB, sel, sort, offset, limit, res)
}

func (self *PostService) Update(dbc *MDBC, sel Selector, data interface{}) {
	dbc.UpdateSet(POST_TAB, sel, data)
}
