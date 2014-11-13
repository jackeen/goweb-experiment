package main

import (
//"log"
//"strconv"
)

type Handler struct {
	TPL  *TPL
	dbc  *MDBC
	post *PostService
	cate *CateService
}

func (self *Handler) Init(dbc *MDBC) {
	self.dbc = dbc
	self.post = &PostService{}
	self.cate = &CateService{}
}

func (self *Handler) Home(req *REQ, res *RES) {

	var postList []Post
	self.post.Select(self.dbc, Selector{}, "id", 0, 10, &postList)

	res.SetHeader("Content-Type", "text/html; charset=utf-8")
	res.State = 200
	res.Response = self.TPL.PostList(postList)
}

func (self *Handler) PostInfo(req *REQ, res *RES) {

	var p Post
	sel := Selector{
		"title": req.PathParm.FileName,
	}

	self.post.SelectOne(self.dbc, sel, &p)

	res.State = 200
	res.Response = self.TPL.Post(p)
}

func (self *Handler) Cate(req *REQ, res *RES) {

	res.State = 200
	res.Response = "post list"
}

func (self *Handler) Tag(req *REQ, res *RES) {

	res.State = 200
	res.Response = "link"
}

func (self *Handler) Date(req *REQ, res *RES) {

	res.State = 200
	res.Response = "link"
}

func (self *Handler) NotFind(req *REQ, res *RES) {

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
