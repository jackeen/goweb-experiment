package main

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Num struct {
	Post int
	Cate int
	Tag  int
	User int
}

type Post struct {
	Id_          bson.ObjectId `bson:"_id"`
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
	Id_     bson.ObjectId `bson:"_id"`
	Id      int
	Content string
	Auth    string
	Email   string
	host    string
	Ip      string
	Reply   []Comment
}

type Cate struct {
	Id_     bson.ObjectId `bson:"_id"`
	Id      int
	Name    string
	Explain string
	Parent  int
}

type Tag struct {
	Name string
}

type User struct {
	Id_   bson.ObjectId `bson:"_id"`
	Id    int
	Name  string
	Pass  string
	Nick  string
	Email string

	CreateTime time.Time
}

type Nav struct {
	Id   string
	Name string
}

type Config struct {
	Host      string
	Copyright string
}
