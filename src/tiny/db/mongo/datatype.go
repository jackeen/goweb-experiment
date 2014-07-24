package mongo

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type IdNum struct {
	Post      int
	Cate      int
	Tag       int
	Comment   int
	User      int
	UserGroup int
}

type Post struct {
	Id_          bson.ObjectId `bson:_id`
	Id           int
	Title        string
	Content      string
	Auth         mgo.DBRef
	Cate         mgo.DBRef
	Tags         mgo.DBRef
	CreateTime   time.Time
	LastEditTime time.Time
	EditState    bool
	AllowComment bool
	Comment      []Comment
}

type Comment struct {
	Id_     bson.ObjectId `bson:_id`
	Id      int
	PostId  int
	Content string
	Auth    string
	Email   string
	host    string
	Ip      string
}

type Cate struct {
	Id_      bson.ObjectId `bson:_id`
	Id       int
	Name     string
	Explain  string
	Children []mgo.DBRef
	Parent   mgo.DBRef
}

type Tag struct {
	Id_     bson.ObjectId `bson:_id`
	Id      int
	Name    string
	Explain string
}

type UserGroup struct {
	Id_ bson.ObjectId `bson:_id`
}

type User struct {
	Id_        bson.ObjectId `bson:_id`
	Id         int
	Name       string
	Pass       string
	Nick       string
	CreateTime time.Time
}

type Config struct {
	Host      string
	Copyright string
}
