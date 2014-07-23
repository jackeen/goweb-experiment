package mongo

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"time"
)

type Tpost struct {
	Title      string
	Content    string
	Auth       mgo.DBRef
	CreateTime time.Time
}

func initPost(db *mgo.Database) {

	//firstTime := time.Now().Format("2006-01-02 15:04:05")

	firstPost := &Tpost{
		/*Title        string
		Content      string
		Auth         mgo.DBRef
		Cate         mgo.DBRef
		Tags         mgo.DBRef
		CreateTime   time.Time
		LastEditTime time.Time
		EditState    bool
		AllowComment bool
		Comment      mgo.DBRef*/

		Title:      "my first post",
		Content:    "hello world!",
		CreateTime: time.Now(),
		//Auth:       mgo.DBRef{"user", "admin"},
	}

	res := &User{}
	err := db.C("user").Find(bson.M{}).One(res)

	ref := mgo.DBRef{
		Collection: "user",
		Id:         res.Name,
	}

	firstPost.Auth = ref

	c := db.C("post")
	err = c.Insert(firstPost)

	if err != nil {
		panic(err)
	}

	log.Println("The post collction is done!")
}

/*
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
}*/

func initUser(db *mgo.Database) {

	root := &User{
		Name:       "admin",
		Nick:       "admin",
		Pass:       "123456",
		CreateTime: time.Now(),
	}

	c := db.C("user")
	err := c.Insert(root)

	if err != nil {
		panic(err)
	}

	log.Println("The user collction is done!")
}

func selPost(db *mgo.Database) {

	res := &Tpost{}

	q := db.C("post").Find(bson.M{})
	err := q.One(res)

	log.Println(res.Auth, err)
}

func Init() {

	log.Println("start tinyblog database ...")

	s, err := mgo.Dial("mongodb://tinyblog:1234@localhost/tinyblog")
	defer s.Close()

	if err != nil {
		panic(err)
	}

	db := s.DB("tinyblog")

	initUser(db)
	//initCate(db)
	initPost(db)

	selPost(db)

}
