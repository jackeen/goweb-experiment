package main

import (
//"encoding/json"
//"log"
)

type Admin struct {
	DBC        *MDBC
	Tpl        *TPL
	StaticHost string
	Session    *Session
	DS         *DataService
}

func (self *Admin) Auth(req *REQ, res *RES) bool {

	c, err := req.R.Cookie("uuid")

	if err == nil {
		if self.Session.IsLogin(c.Value) {
			return true
		} else {
			GotoLogin(req, res)
			return false
		}
	} else {
		GotoLogin(req, res)
		return false
	}
}

func (self *Admin) Router(req *REQ, res *RES) {

	s := self.Auth(req, res)
	if !s {
		return
	}

	switch req.PathParm.FileName {
	case "home":
		self.home(req, res)
	case "addpost":
		self.addPost(req, res)
	case "postlist":
		self.postList(req, res)
	default:
		self.NotFound(req, res)
	}

}

func (self *Admin) home(req *REQ, res *RES) {

	d := map[string]interface{}{
		"PageTitle": "manager home",
	}

	res.Response = self.Tpl.Parse("home", d)
}

func (self *Admin) addPost(req *REQ, res *RES) {

	d := map[string]interface{}{
		"PageTitle": "add post",
	}

	res.Response = self.Tpl.Parse("add_post", d)
}

func (self *Admin) postList(req *REQ, res *RES) {

	d := map[string]interface{}{
		"PageTitle": "post list",
	}

	res.Response = self.Tpl.Parse("post_list", d)
}

func (self *Admin) NotFound(req *REQ, res *RES) {
	res.State = 404
	res.Response = "none"
}
