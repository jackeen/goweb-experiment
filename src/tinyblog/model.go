package main

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Post struct {
	Id_          bson.ObjectId `bson:"_id"`
	Title        string
	Content      string
	Author       string
	Cate         string
	Tags         []string
	CreateTime   time.Time
	EditTime     time.Time
	IsDraft      bool
	AllowComment bool
	Comment      []Comment
}

type Comment struct {
	Id_        bson.ObjectId `bson:"_id"`
	Content    string
	Auth       string
	Ip         string
	IsDisplay  bool
	CreateTime time.Time
}

type Cate struct {
	Name    string
	Level   int
	Parent  string
	Explain string
}

type Tag struct {
	Name string
}

type Config struct {
	HostName   string
	Copyright  string
	Explain    string
	CreateTime time.Time
	EditTime   time.Time
}

type User struct {
	Id_        bson.ObjectId `bson:"_id"`
	Name       string
	Pass       string
	Nick       string
	Email      string
	Face       string
	PowerCode  int
	CreateTime time.Time
}
