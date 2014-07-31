package main

import (
	db "db/mongo"
	"log"
	"time"
)

func FindPost() {

}

func InsertPost(title string, content string) {

	data := &Post{
		Id:         1,
		Title:      title,
		Content:    content,
		Auth:       "admin",
		Cate:       -1,
		Tags:       "",
		CreateTime: time.Now(),
	}

	db.InitDBC(DBConfig)

	db.Insert(DBTable["Post"], data)

}

func main() {

	InsertPost("title a", "content")
	log.Println("aa")
}
