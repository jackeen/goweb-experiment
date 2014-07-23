package mongo

import (
	"labix.org/v2/mgo"
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
	Id           int
	Title        string
	Content      string
	Auth         int
	Cate         int
	Tags         int
	CreateTime   time.Time
	LastEditTime time.Time
	EditState    bool
	AllowComment bool
	Comment      []interface{}
}

type Comment struct {
	Id      int
	PostId  int
	Content string
	Auth    string
	Email   string
	host    string
	Ip      string
}

type Cate struct {
	Id       int
	Name     string
	Explain  string
	Children mgo.DBRef
	Parent   int
}

type Tag struct {
	Id      int
	Name    string
	Explain string
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

type Config struct {
	Host      string
	Copyright string
}
