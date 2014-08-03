package main

import (
	"log"
)

const (
	POST_TAB   = "post"
	CATE_TAB   = "cate"
	TAG_TAB    = "tag"
	NAV_TAB    = "nav"
	CONFIG_TAB = "config"
)

var dbc *MDBC

func initDB(host string, user string, pass string, name string) {
	dbc = &MDBC{
		Host: host,
		User: user,
		Pass: pass,
		Name: name,
	}
	dbc.Init()
}

func main() {

	var postList []Post

	initDB("localhost", "tinyblog", "1234", "tinyblog")
	postSer := &PostService{}

	//postSer.Insert(dbc, "title", "content")
	//postSer.Find(dbc, nil, "id", 0, 2, &postList)
	postSer.Update(dbc, BSON{"id": 1}, BSON{"title": "hahahaha"})
	log.Println(postList)

}
