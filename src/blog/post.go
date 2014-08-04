package main

import (
	//"log"
	//"reflect"
	//"errors"
	"labix.org/v2/mgo/bson"
	"time"
)

func increaseNumId(dbc *MDBC, colName string, i int) *IdNum {

	res := &IdNum{}
	dbc.UpdateInc(NUM_TAB, nil, colName, i)
	dbc.Select(NUM_TAB, nil, "", 0, 1, res)
	return res
}

type PostService struct {
}

func (self *PostService) Insert(dbc *MDBC, title string, content string) {

	data := &Post{
		Id_:        bson.NewObjectId(),
		Id:         increaseNumId(dbc, "post", 1).Post,
		Title:      title,
		Content:    content,
		Auth:       "admin",
		Cate:       -1,
		Tags:       "",
		CreateTime: time.Now(),
	}
	dbc.Insert(POST_TAB, data)
}

func (self *PostService) Find(dbc *MDBC, selector BSON, sort string, offset int, limit int, res *[]Post) {

	if sort == "" {
		sort = "id"
	}

	dbc.Select(POST_TAB, nil, sort, offset, limit, res)
}

func (self *PostService) Update(dbc *MDBC, selector BSON, data interface{}) {
	dbc.UpdateSet(POST_TAB, selector, data)
}
