package main

import (
	"time"
)

type IdNum struct {
	Post int
	Cate int
	Tag  int
}

type Post struct {
	Id_          MongoId `bosn:"_id"`
	Id           int
	Title        string "title"
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
	Id_     MongoId `bosn:"_id"`
	Id      int
	Content string
	Auth    string
	Email   string
	host    string
	Ip      string
	Reply   []Comment
}

type Cate struct {
	Id_     MongoId `bosn:"_id"`
	Id      int
	Name    string
	Explain string
	Parent  int
}

type Tag struct {
	Id   int
	Name string
}

/*
type UserGroup struct {
}

type User struct {
	Id         int
	Name       string
	Pass       string
	Nick       string
	CreateTime time.Time
}
*/

type Nav struct {
	Id   string
	Name string
}

type Config struct {
	Host      string
	Copyright string
}
