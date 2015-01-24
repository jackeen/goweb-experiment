package main

import (
//"log"
//"strconv"
)

type Handler struct {
	Tpl        *TPL
	StaticHost string
	dbc        *MDBC
	post       *PostService
	cate       *CateService
}

func (self *Handler) Init(dbc *MDBC) {
	self.dbc = dbc
	self.post = &PostService{
		DBC: dbc,
	}
	self.cate = &CateService{
		DBC: dbc,
	}
}

func (self *Handler) Index(req *REQ, res *RES) {

	var postList []Post

	selData := &SelectData{
		Sort:  "-createtime",
		Limit: 10,
		Res:   &postList,
	}
	self.post.Select(selData)

	d := map[string]interface{}{
		"PageTitle":  "home",
		"StaticHost": self.StaticHost,
		"PostList":   postList,
	}

	res.Response = self.Tpl.Parse("index", d)
}

func (self *Handler) Post(req *REQ, res *RES) {

	var p Post
	sel := BSONM{
		"title": req.PathParm.FileName,
	}

	self.post.SelectOne(sel, &p)

	res.Response = self.Tpl.Parse("post", p)
}

func (self *Handler) Cate(req *REQ, res *RES) {

	res.Response = "post list"
}

func (self *Handler) Tag(req *REQ, res *RES) {

	res.Response = "link"
}

func (self *Handler) Date(req *REQ, res *RES) {

	res.Response = "Date"
}

func (self *Handler) Entry(req *REQ, res *RES) {

	d := map[string]interface{}{
		"PageTitle":  "~~login~~",
		"StaticHost": self.StaticHost,
	}

	res.Response = self.Tpl.Parse("login", d)
}

func (self *Handler) NotFind(req *REQ, res *RES) {

	res.State = 404
	res.Response = "none"
}
