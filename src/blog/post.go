package main

import (
	"log"
	//"reflect"
	"time"
)

func insertPost(dbc *MDBC, title string, content string) {

	data := &Post{
		Id:         1,
		Title:      title,
		Content:    content,
		Auth:       "admin",
		Cate:       -1,
		Tags:       "",
		CreateTime: time.Now(),
	}

	c := dbc.DB.C(POST_TAB)
	err := c.Insert(data)
	if err != nil {
		log.Println(err)
	}
}

func findPost(dbc *MDBC, sel Selector, res *[]Post) {

	c := dbc.DB.C(POST_TAB)
	f := c.Find(sel)

	err := f.Limit(3).All(res)
	if err != nil {
		log.Println(err)
	}
}
