package main

import (
	"labix.org/v2/mgo/bson"
)

type Post struct {
	Id_          bson.ObjectId `bson:"_id" json:"id"`
	Title        string        `bson:"title" json:"title"`
	Content      string        `bson:"content" json:"content"`
	Author       string        `bson:"author" json:"author"`
	Cate         string        `bson:"cate" json:"cate"`
	Tags         []string      `bson:"tags" json:"tags"`
	CreateTime   TimeData      `bson:"createtime" json:"createTime"`
	EditTime     TimeData      `bson:"edittime" json:"editTime"`
	IsDraft      bool          `bson:"isdraft" json:"isDraft"`
	IsDiscard    bool          `bson:"isdiscard" json:"isDiscard"`
	AllowComment bool          `bson:"allowcomment" json:"allowComment"`
	Comment      []Comment     `bson:"comment" json:"comment"`
}

type Comment struct {
	Id_        bson.ObjectId `bson:"_id" json:"id"`
	Content    string        `bson:"content" json:"content"`
	Auth       string        `bson:"auth" json:"auth"`
	Ip         string        `bson:"ip" json:"ip"`
	IsDisplay  bool          `bson:"isdisplay" json:"isDisplay"`
	CreateTime TimeData      `bson:"createtime" json:"createTime"`
	EditTIme   TimeData      `bson:"edittime" json:"editTime"`
	Reply      string        `bson:"reply" json:"reply"`
}

type Cate struct {
	Name     string   `bson:"name" json:"name"`
	Children []string `bson:"children" json:"children"`
	Parent   string   `bson:"parent" json:"parent"`
	PLink    string   `bson:"plink" json:"plink"`
	//Level    int      `bson:"level" json:"level"`
}

type Tag struct {
	Name string `bson:"name" json:"name"`
}

type Image struct {
	Id_         bson.ObjectId `bson:"_id" json:"id"`
	FileName    string        `bson:"filename" json:"fileName"`
	Name        string        `bson:"name" json:"name"`
	ContentName string        `bson:"contentname" json:"contentName"`
	Size        string        `bson:"size" json:"size"`
	Cate        string        `bson:"cate" json:"cate"`
	Author      string        `bson:"author" json:"author"`
	CreateTime  TimeData      `bson:"createtime" json:"createTime"`
	EditTime    TimeData      `bson:"edittime" json:"editTime"`
}

type ImageMeta struct {
	ContentName string
	Name        string
}

type ImageCate struct {
	Id_        bson.ObjectId `bson:"_id" json:"id"`
	Name       string        `bson:"name" json:"name"`
	Explain    string        `bson:"explain" json:"explain"`
	CreateTime TimeData      `bson:"createtime" json:"createTime"`
	EditTime   TimeData      `bson:"edittime" json:"editTime"`
}

type Config struct {
	HostName   string   `bson:"hostname" json:"hostName"`
	Copyright  string   `bson:"copyright" json:"copyRight"`
	Explain    string   `bson:"explain" json:"explain"`
	CreateTime TimeData `bson:"createtime" json:"createTime"`
	EditTime   TimeData `bson:"edittime" json:"editTime"`
}

type User struct {
	Id_        bson.ObjectId `bson:"_id" json:"id"`
	Name       string        `bson:"name" json:"name"`
	Pass       string        `bson:"pass" json:"pass"`
	Nick       string        `bson:"nick" json:"nick"`
	Email      string        `bson:"email" json:"email"`
	Face       string        `bson:"face" json:"face"`
	CreateTime TimeData      `bson:"createtime" json:"createTime"`
	ditTime    TimeData      `bson:"edittime" json:"editTime"`
	Active     bool          `bson:"active" json:"active"`
	Group      string        `bson:"group" json:"group"`
}
