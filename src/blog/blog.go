package main

import (
	"flag"
	"log"
)

const (
	NUM_TAB    = "num"
	POST_TAB   = "post"
	CATE_TAB   = "cate"
	USER_TAB   = "user"
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

	//postSer.Insert(dbc, "hi world", "this is a", 1, -1, []string{},false,false)
	postSer.Select(dbc, Selector{"id": *pid}, "id", 0, 2, &postList)
	//postSer.Update(dbc, *pid, Selector{"title": "hahahaha"})

	log.Println(*pid, postList)

	postSer.InsertComment(dbc, *pid, -1, "my comments")
	//postSer.deleteComment(dbc, *pid, 0)

	//userSer := &UserService{}
	//userSer.Insert(dbc, "admin", "1234", "firstuser", "")

}
