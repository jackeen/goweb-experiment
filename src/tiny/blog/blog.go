package main

import (
	"fmt"
	db "tiny/db/mongo"
)

func main() {

	dbc := &db.DBConfig{
		DBHost: "localhost",
		DBUser: "tinyblog",
		DBPass: "1234",
		DBName: "tinyblog",
	}

	tableName := &db.TabName{
		Post:   "post",
		Cate:   "cate",
		Tag:    "tag",
		Nav:    "nav",
		Config: "config",
	}

	postService := &db.PostService{
		DBC: dbc,
		Tab: tableName,
	}

	postService.Insert("aa", "this is a")

	fmt.Println("haha")
}
