package main

import (
	//"log"
	"html/template"
	"strconv"
)

type WriteContent struct {
	S string
}

func (self *WriteContent) Write(p []byte) (n int, err error) {
	self.S += string(p)
	return 0, nil
}

type Handler struct {
	TempLateDir string
	dbc         *MDBC
	post        *PostService
	cate        *CateService
}

func (self *Handler) Init(dbc *MDBC) {
	self.dbc = dbc
	self.post = &PostService{}
	self.cate = &CateService{}
}

func (self *Handler) Home(req *HTTPServerReq, res *HTTPServerRes) {

	var postList []Post
	self.post.Select(self.dbc, Selector{}, "id", 0, 10, &postList)

	restring := ""

	for i, v := range postList {
		restring += (strconv.Itoa(i) + ": " + v.Title + "\n")
	}

	res.State = 200
	res.Response = restring
}

func (self *Handler) PostInfo(req *HTTPServerReq, res *HTTPServerRes) {

	var p Post
	sel := Selector{
		"title": "a",
	}

	self.post.SelectOne(self.dbc, sel, &p)

	res.State = 200
	res.Response = "post info"
}

func (self *Handler) Cate(req *HTTPServerReq, res *HTTPServerRes) {

	res.State = 200
	res.Response = "post list"
}

func (self *Handler) Tag(req *HTTPServerReq, res *HTTPServerRes) {

	res.State = 200
	res.Response = "link"
}

func (self *Handler) Date(req *HTTPServerReq, res *HTTPServerRes) {

	res.State = 200
	res.Response = "link"
}

func (self *Handler) NotFind(req *HTTPServerReq, res *HTTPServerRes) {

	res.State = 404
	res.Response = "none"
}

//var postList []Post

//postServe.Insert(dbc, "hi world", "this is a", 1, -1, []string{}, false, false)
//postServe.Select(dbc, Selector{"id": *pid}, "id", 0, 2, &postList)
//postServe.Update(dbc, *pid, Selector{"title": "hahahaha"})
//log.Println(*pid, postList)

//postServe.InsertComment(dbc, *pid, -1, "my comments")
//postServe.deleteComment(dbc, *pid, 0)

//userSer := &UserService{}
//userSer.Insert(dbc, "admin", "1234", "firstuser", "", 0)

//numid := &NumService{}
//numid.Init(dbc)

//cateSer := &CateService{}
//cate := &Cate{}
//cateSer.Inert(dbc, "js", "javascript", 0)
//cateSer.Update(dbc, *cid, Selector{"name": "ajax"})
//cateSer.Select(dbc, *cid, cate)
//log.Println(cate)
//cateSer.Delete(dbc, *cid)
