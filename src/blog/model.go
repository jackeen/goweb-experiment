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
	Author       int
	Cate         int
	Tags         []string
	CreateTime   time.Time
	LastEditTime time.Time
	Draft        bool
	AllowComment bool
	CommentNum   int
	CommentIncId int
	Comment      []Comment
}

type Comment struct {
	Id_        bson.ObjectId `bson:"_id"`
	Id         int
	Content    string
	Auth       string
	Email      string
	host       string
	Ip         string
	Display    bool
	ReplyId    int
	CreateTime time.Time
}

type Cate struct {
	Id_      bson.ObjectId `bson:"_id"`
	Id       int
	Name     string
	Explain  string
	ParentId int
}

type Tag struct {
	Name string
}

type User struct {
	Id_        bson.ObjectId `bson:"_id"`
	Id         int
	Name       string
	Pass       string
	Nick       string
	Email      string
	Power      int
	CreateTime time.Time
}

type Nav struct {
	Id         int
	Name       string
	CreateTime time.Time
	EditTime   time.Time
}

type Config struct {
	Host       string
	Copyright  string
	CreateTime time.Time
	EditTime   time.Time
}
