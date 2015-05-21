package main

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Post struct {
	Id_          bson.ObjectId `bson:"_id" json:"id"`
	Title        string        `bson:"title" json:"title"`
	Content      string        `bson:"content" json:"content"`
	Author       string        `bson:"author" json:"author"`
	Cate         string        `bson:"cate" json:"cate"`
	Tags         []string      `bson:"tags" json:"tags"`
	CreateTime   time.Time     `bson:"createtime" json:"createTime"`
	EditTime     time.Time     `bson:"edittime" json:"editTime"`
	IsDraft      bool          `bson:"isdraft" json:"isDraft"`
	AllowComment bool          `bson:"allowcomment" json:"allowComment"`
	Comment      []Comment     `bson:"comment" json:"comment"`
}

type Comment struct {
	Id_        bson.ObjectId `bson:"_id" json:"id"`
	Content    string        `bson:"content" json:"content"`
	Auth       string        `bson:"auth" json:"auth"`
	Ip         string        `bson:"ip" json:"ip"`
	IsDisplay  bool          `bson:"isdisplay" json:"isDisplay"`
	CreateTime time.Time     `bson:"createtime" json:"createTime"`
	EditTIme   time.Time     `bson:"edittime" json:"editTime"`
	Reply      string        `bson:"reply" json:"reply"`
}

type Cate struct {
	Name    string `bson:"name"`
	Level   int    `bson:"level"`
	Parent  string `bson:"parent"`
	Explain string `bson:"explain"`
}

type Tag struct {
	Name string `bson:"name" json:"name"`
}

type Config struct {
	HostName   string    `bson:"hostname" json:"hostName"`
	Copyright  string    `bson:"copyright" json:"copyRight"`
	Explain    string    `bson:"explain" json:"explain"`
	CreateTime time.Time `bson:"createtime" json:"createTime"`
	EditTime   time.Time `bson:"edittime" json:"editTime"`
}

type User struct {
	Id_        bson.ObjectId `bson:"_id" json:"id"`
	Name       string        `bson:"name" json:"name"`
	Pass       string        `bson:"pass" json:"pass"`
	Nick       string        `bson:"nick" json:"nick"`
	Email      string        `bson:"email" json:"email"`
	Face       string        `bson:"face" json:"face"`
	PowerCode  int           `bson:"powercode" json:"powerCode"`
	CreateTime time.Time     `bson:"createtime" json:"createTime"`
}
