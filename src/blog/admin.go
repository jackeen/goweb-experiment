package main

import (
//"log"
)

type LoginComplete struct {
	Name string
}

type Admin struct {
	DBC    *MDBC
	TPLDIR string
	tpl    *TPL
}

func (self *Admin) Init() {
	self.tpl = &TPL{
		TmpDir: self.TPLDIR,
	}
}

func (self *Admin) Router(req *REQ, res *RES) {

	switch req.PathParm.FileName {
	case "entry":
		self.entry(req, res)
	default:
		self.notFound(req, res)
	}

}

//login and logout
func (self *Admin) entry(req *REQ, res *RES) {

	loginPage := func() {

		res.State = 200
		res.Response = self.tpl.Login(nil)
	}

	switch req.GetUrlOneValue("action") {
	case "login":
		self.login(req, res)
	case "logout":
		self.logout(req, res)
	default:
		loginPage()
	}
}

func (self *Admin) login(req *REQ, res *RES) {

	u := req.GetFormValue("user")
	p := req.GetFormValue("pass")

	user := &User{}

	uc := &UserService{}
	uc.LoginSelect(self.DBC, u, p, user)

	res.State = 200
	res.Response = self.tpl.LoginComplete(user)
}

func (self *Admin) logout(req *REQ, res *RES) {

}

//post
func (self *Admin) post(req *REQ, res *RES) {

}

func (self *Admin) notFound(req *REQ, res *RES) {
	res.State = 301
	res.SetHeader("Location", "/")
}
