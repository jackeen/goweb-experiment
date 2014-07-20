package mongo

import (
	"labix.org/v2/mgo"
	"log"
	"time"
)

func initPost(db *mgo.Database) {

	firstTime := time.Now().Format("2006-01-02 15:04:05")

	firstPost := &Post{
		Id:      "2",
		Title:   "my first post",
		Content: "hello world!",
		Auth:    "tiny",
		AddDate: firstTime,
	}

	c := db.C("post")
	err := c.Insert(firstPost)

	if err != nil {
		panic(err)
	}

	log.Println("The post collction is done!")
}

func initCate(db *mgo.Database) {

	rootCate := &Cate{
		Id:       "0",
		Name:     "root",
		Explain:  "",
		Children: []string{},
		Parent:   "-1",
	}

	c := db.C("cate")
	err := c.Insert(rootCate)

	if err != nil {
		panic(err)
	}

	log.Println("The cate collction is done!")
}

func Init() {

	log.Println("start tinyblog database ...")

	s, err := mgo.Dial("mongodb://tinyblog:1234@localhost/tinyblog")
	defer s.Close()

	if err != nil {
		panic(err)
	}

	db := s.DB("tinyblog")

	initPost(db)
	initCate(db)

}
