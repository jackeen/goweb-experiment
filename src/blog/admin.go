package main

import (
//"encoding/json"
//"log"
)

type Admin struct {
	DBC        *MDBC
	TPL        *AdminTPL
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
	default:
		self.NotFound(req, res)
	}

}

func (self *Admin) home(req *REQ, res *RES) {
	d := &AdminHomeData{
		PageTitle: "manager home",
	}
	res.State = 200
	res.Response = self.TPL.Home(d)
}

func (self *Admin) NotFound(req *REQ, res *RES) {
	res.State = 404
	res.Response = "none"
}
