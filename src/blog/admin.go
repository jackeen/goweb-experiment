package main

import (
//"log"
)

type LoginComplete struct {
	Name string
}

type Admin struct {
	DBC    *MDBC
	tpl    *TPL
	TplDir string
}

func (self *Admin) Init() {
	self.tpl = &TPL{
		TmpDir: self.TplDir,
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

	switch req.GetUrlOneValue("action") {
	case "login":
		self.login(req, res)
	case "logout":
		self.logout(req, res)
	default:
		res.State = 200
		res.Response = self.tpl.Login(nil)
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
	switch req.GetUrlOneValue("action") {
	case "add":
		self.addPost(req, res)
	case "edit":
		self.editPost(req, res)
	case "delete":
		self.deletePost(req, res)
	default:
		res.State = 404
		res.Response = "404"
	}
}

func (self *Admin) addPost(req *REQ, res *RES) {

}

func (self *Admin) editPost(req *REQ, res *RES) {

}

func (self *Admin) deletePost(req *REQ, res *RES) {

}

func (self *Admin) notFound(req *REQ, res *RES) {
	res.State = 301
	res.SetHeader("Location", "/")
}
