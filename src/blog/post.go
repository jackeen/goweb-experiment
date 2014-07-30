package main

import (
	db "db/mongo"
	//"log"
	"time"
)

func InsertPost(dbc *db.Config, tab *db.TabName, title string, content string) {

	data := &Post{
		Id:         1,
		Title:      title,
		Content:    content,
		Auth:       "admin",
		Cate:       -1,
		Tags:       "",
		CreateTime: time.Now(),
	}

	db.Execute(dbc, tab.Post, db.Insert, data)

}

func main() {

	dbc := &db.Config{
		Host: "localhost",
		User: "tinyblog",
		Pass: "1234",
		Name: "tinyblog",
	}
	tab := db.GetTabName()

	InsertPost(dbc, tab, "title a", "content")

}
