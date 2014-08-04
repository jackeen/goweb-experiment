package main

import (
	//"log"
	//"reflect"
	//"errors"
	"time"
)

type PostService struct {
}

func (self *PostService) Insert(dbc *MDBC, title string, content string) {

	data := &Post{
		Id_:        dbc.GetMongoId(),
		Id:         1,
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
