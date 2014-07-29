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

	db.PostInsert(dbc, "aa", "this is a")

	fmt.Println("haha")
}
