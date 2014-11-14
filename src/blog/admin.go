package main

import (
	//"log"
	"encoding/json"
)

var (
	dbc           *MDBC
	tpl           *AdminTPL
	entryMod      *entry
	staticRootURL string
)

func adminNotFoundPage(req *REQ, res *RES) {

}

func adminNotFoundJson(req *REQ, res *RES) {
	m := map[string]interface{}{
		"success": false,
		"code":    404,
		"message": "not found",
	}
	bytes, _ := json.Marshal(m)
	res.Response = string(bytes)
}

type Admin struct {
	DBC           *MDBC
	TPL           *AdminTPL
	StaticRootURL string
}

func (self *Admin) Init() {

	dbc = self.DBC
	tpl = self.TPL
	staticRootURL = self.StaticRootURL

	entryMod = &entry{}
}

func (self *Admin) Router(req *REQ, res *RES) {

	switch req.PathParm.FileName {
	case "entry":
		entryMod.Route(req, res)
	default:
		adminNotFoundPage(req, res)
	}

}

//
type entry struct{}

func (self *entry) Route(req *REQ, res *RES) {

	switch req.GetUrlOneValue("serve") {
	case "":
		self.loginPage(req, res)
	case "login":
		self.loginServe(req, res)
	case "logout":
		self.logoutServe(req, res)
	default:
		adminNotFoundJson(req, res)
	}
}

func (self *entry) loginPage(req *REQ, res *RES) {
	res.State = 200
	res.Response = tpl.Login(nil)
}

func (self *entry) loginServe(req *REQ, res *RES) {

	u := req.GetFormValue("user")
	p := req.GetFormValue("pass")

	user := &User{}

	uc := &UserService{}
	uc.LoginSelect(dbc, u, p, user)

	res.State = 200
	res.Response = tpl.LoginComplete(user)
}

func (self *entry) logoutServe(req *REQ, res *RES) {

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
