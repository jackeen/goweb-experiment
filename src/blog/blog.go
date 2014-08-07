package main

import (
	"flag"
	"log"
)

const (
	NUM_TAB    = "num"
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

	pid := flag.Int("pid", 0, "post select id")
	flag.Parse()

	var postList []Post

	initDB("localhost", "tinyblog", "1234", "tinyblog")
	postSer := &PostService{}

	//postSer.Insert(dbc, "title", "content")
	postSer.Find(dbc, Selector{"id": *pid}, "id", 0, 2, &postList)
	//postSer.Update(dbc, Selector{"id": 1}, Selector{"title": "hahahaha"})
	log.Println(*pid, postList)

}
