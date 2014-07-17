package mongo

import (
	"labix.org/v2/mgo"
	"log"
	"time"
)

func Init() {

	log.Println("start tinyblog database ...")

	s, err := mgo.Dial("mongodb://tinyblog:1234@localhost/tinyblog")
	defer s.Close()

	if err != nil {
		panic(err)
	}

	firstTime := time.Now().Format("2006-01-02 15:04:05")

	firstPost := &Post{
		id:      "1",
		title:   "my first post",
		content: "hello world!",
		auth:    "tiny",
		addDate: firstTime,
	}

	c := s.DB("tinyblog").C("post")
	err = c.Insert(firstPost)

	if err != nil {
		panic(err)
	}

	log.Println(err)
}
