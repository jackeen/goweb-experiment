package mongo

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"time"
)

type Tpost struct {
	Id_        bson.ObjectId `bson:"_id"`
	Title      string
	Content    string
	Auth       mgo.DBRef
	CreateTime time.Time
}

type Tuser struct {
	Id_  bson.ObjectId `bson:"_id"`
	Name string
	Pass string
	Nick string
}

func initPost(db *mgo.Database) {

	//firstTime := time.Now().Format("2006-01-02 15:04:05")

	firstPost := &Tpost{
		Id_:        bson.NewObjectId(),
		Title:      "my first post",
		Content:    "hello world!",
		CreateTime: time.Now(),
	}

	res := &Tuser{}
	err := db.C("user").Find(bson.M{}).One(res)

	ref := mgo.DBRef{
		Collection: "user",
		Id:         res.Id_,
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

	root := &Tuser{
		Id_:  bson.NewObjectId(),
		Name: "admin",
		Pass: "123456",
		Nick: "admin",
	}

	c := db.C("user")
	err := c.Insert(root)

	if err != nil {
		panic(err)
	}

	log.Println("The user collction is done!")
}

func selPost(db *mgo.Database) {

	var posts []Tpost

	q := db.C("post").Find(nil)

	err := q.All(&posts)

	usrRef := posts[0].Auth

	rq := db.FindRef(&usrRef)
	refUser := &Tuser{}
	err = rq.One(refUser)

	log.Println(refUser, err)
}

func init() {

	log.Println("start tinyblog database ...")

	s, err := mgo.Dial("mongodb://tinyblog:1234@localhost/tinyblog")
	defer s.Close()

	if err != nil {
		panic(err)
	}

	db := s.DB("tinyblog")

	//initUser(db)

	//initPost(db)

	selPost(db)

}
