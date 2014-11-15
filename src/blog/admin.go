package main

import (
	//"log"
	"encoding/json"
)

type Admin struct {
	DBC           *MDBC
	TPL           *AdminTPL
	StaticRootURL string
	entry         *EntryPage
}

func (self *Admin) Init() {

	self.entry = &EntryPage{
		Parent: self,
	}
}

func (self *Admin) Router(req *REQ, res *RES) {

	switch req.PathParm.FileName {
	case "entry":
		self.entry.Route(req, res)
	default:
		self.AdminNotFoundPage(req, res)
	}

}

func (self *Admin) AdminNotFoundPage(req *REQ, res *RES) {
	res.State = 200
	res.Response = "404 page"
}

func (self *Admin) AdminNotFoundJson(req *REQ, res *RES) {
	m := map[string]interface{}{
		"success": false,
		"code":    404,
		"message": "not found",
	}
	bytes, _ := json.Marshal(m)
	res.State = 200
	res.Response = string(bytes)
}

//
type EntryPage struct {
	Parent *Admin
}

func (self *EntryPage) Route(req *REQ, res *RES) {

	switch req.GetUrlOneValue("serve") {
	case "":
		self.loginPage(req, res)
	case "login":
		self.loginServe(req, res)
	case "logout":
		self.logoutServe(req, res)
	default:
		self.Parent.AdminNotFoundJson(req, res)
	}
}

func (self *EntryPage) loginPage(req *REQ, res *RES) {
	res.State = 200
	res.Response = self.Parent.TPL.Login(nil)
}

func (self *EntryPage) loginServe(req *REQ, res *RES) {

	u := req.GetFormValue("user")
	p := req.GetFormValue("pass")

	user := &User{}

	uc := &UserService{}
	uc.LoginSelect(self.Parent.DBC, u, p, user)

	res.State = 200
	res.Response = self.Parent.TPL.LoginComplete(user)
}

func (self *EntryPage) logoutServe(req *REQ, res *RES) {

}

/*

//login and logout
func (self *Admin) entry(req *REQ, res *RES) {

	switch req.GetUrlOneValue("action") {
	case "login":
		self.login(req, res)
	case "logout":
		self.logout(req, res)
	default:
		res.State = 200
		res.Response = self.TPL.Login(nil)
	}
}

func (self *Admin) login(req *REQ, res *RES) {

	u := req.GetFormValue("user")
	p := req.GetFormValue("pass")

	user := &User{}

	uc := &UserService{}
	uc.LoginSelect(self.DBC, u, p, user)

	res.State = 200
	res.Response = self.TPL.LoginComplete(user)
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


*/
