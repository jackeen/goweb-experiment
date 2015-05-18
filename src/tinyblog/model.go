package main

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Post struct {
	Id_          bson.ObjectId `bson:"_id"`
	Title        string        `bson:"title"`
	Content      string        `bson:"content"`
	Author       string        `bson:"author"`
	Cate         string        `bson:"cate"`
	Tags         []string      `bson:"tags"`
	CreateTime   time.Time     `bson:"createtime"`
	EditTime     time.Time     `bson:"edittime"`
	IsDraft      bool          `bson:"isdraft"`
	AllowComment bool          `bson:"allowcomment"`
	Comment      []Comment     `bson:"comment"`
}

type Comment struct {
	Id_        bson.ObjectId `bson:"_id"`
	Content    string        `bson:"content"`
	Auth       string        `bson:"auth"`
	Ip         string        `bson:"ip"`
	IsDisplay  bool          `bson:"isdisplay"`
	CreateTime time.Time     `bson:"createtime"`
}

type Cate struct {
	Name    string `bson:"name"`
	Level   int    `bson:"level"`
	Parent  string `bson:"parent"`
	Explain string `bson:"explain"`
}

type Tag struct {
	Name string `bson:"name"`
}

type Config struct {
	HostName   string    `bson:"hostname"`
	Copyright  string    `bson:"copyright"`
	Explain    string    `bson:"explain"`
	CreateTime time.Time `bson:"createtime"`
	EditTime   time.Time `bson:"edittime"`
}

type User struct {
	Id_        bson.ObjectId `bson:"_id"`
	Name       string        `bson:"name"`
	Pass       string        `bson:"pass"`
	Nick       string        `bson:"nick"`
	Email      string        `bson:"email"`
	Face       string        `bson:"face"`
	PowerCode  int           `bson:"powercode"`
	CreateTime time.Time     `bson:"createtime"`
}
