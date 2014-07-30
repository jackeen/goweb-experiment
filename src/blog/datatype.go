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
	User      int
	UserGroup int
}

type Post struct {
	//Id_          bson.ObjectId `bson:_id`
	Id           int
	Title        string
	Content      string
	Auth         string
	Cate         int
	Tags         string
	CreateTime   time.Time
	LastEditTime time.Time
	EditState    bool
	AllowComment bool
	TotalComment int
	Comment      []Comment
}

type Comment struct {
	//Id_     bson.ObjectId `bson:_id`
	Id      int
	Content string
	Auth    string
	Email   string
	host    string
	Ip      string
	Reply   []Comment
}

type Cate struct {
	//Id_     bson.ObjectId `bson:_id`
	Id      int
	Name    string
	Explain string
	Parent  mgo.DBRef
}

type Tag struct {
	Id   int
	Name string
}

type UserGroup struct {
}

type User struct {
	Id         int
	Name       string
	Pass       string
	Nick       string
	CreateTime time.Time
}

/* cms config */

type Nav struct {
	Id   string
	Name string
}

type Config struct {
	Host      string
	Copyright string
}
