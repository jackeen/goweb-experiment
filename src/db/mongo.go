package db

import (
	"labix.org/v2/mgo"
	"log"
	"time"
)

type Post struct {
	Id      string
	Title   string
	Content string
	Auth    string
	AddDate string
}

func InitDB() {

	log.Println("start tinyblog database ...")

	s, err := mgo.Dial("mongodb://tinyblog:1234@localhost/tinyblog")
	defer s.Close()

	if err != nil {
		panic(err)
	}

	firstTime := time.Now().Format("2006-01-02 15:04:05")

	firstPost := &Post{
		Id:      "1",
		Title:   "my first post",
		Content: "hello world!",
		Auth:    "tiny",
		AddDate: firstTime,
	}

	c := s.DB("tinyblog").C("post")
	err = c.Insert(firstPost)

	if err != nil {
		panic(err)
	}

	log.Println(err)
}
