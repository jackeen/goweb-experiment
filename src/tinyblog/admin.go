package main

import (
//"encoding/json"
//"log"
)

type Admin struct {
	DBC        *MDBC
	Tpl        *TPL
	StaticHost string
	session    *Session
}

func (self *Admin) Init(s *Session) {
	self.session = s
}

func (self *Admin) Router(req *REQ, res *RES) {

	switch req.PathParm.FileName {
	case "home":
		self.home(req, res)
	case "addpost":
		self.addPost(req, res)
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
		"PageTitle":  "add post",
		"StaticHost": self.StaticHost,
	}

	res.Response = self.Tpl.Parse("addPost", d)
}

func (self *Admin) NotFound(req *REQ, res *RES) {
	res.State = 404
	res.Response = "none"
}
