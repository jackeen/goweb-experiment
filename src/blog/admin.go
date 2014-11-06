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

	action := req.Query["action"]

	loginPage := func() {
		tpl := &TPL{
			TmpDir: self.TPLDIR,
		}
		res.State = 200
		res.Response = tpl.Login(nil)
	}

	if len(action) == 1 {

		switch action[0] {
		case "login":
			self.login(req, res)
		case "logout":
			self.logout(req, res)
		default:
			loginPage()
		}

	} else {
		loginPage()
	}
}

func (self *Admin) login(req *REQ, res *RES) {

	res.State = 200
	res.Response = "x"
}

func (self *Admin) logout(req *REQ, res *RES) {

}

//post
func (self *Admin) addPost(req *REQ, res *RES) {

}

func (self *Admin) notFound(req *REQ, res *RES) {
	res.State = 301
	res.Headers = map[string]string{
		"Location": "/",
	}
}
