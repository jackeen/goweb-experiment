package main

import (
	"log"
)

type LoginComplete struct {
	Name string
}

type Admin struct {
	DBC    *MDBC
	TPLDIR string
}

func (self *Admin) Action(req *HTTPServerReq, res *HTTPServerRes) {

	switch req.PathParm.FileName {
	case "login":
		self.login(req, res)
	default:
		self.notFound(req, res)
	}

}

func (self *Admin) login(req *HTTPServerReq, res *HTTPServerRes) {
	/*
		userName := req.Req.FormValue("user")
		user := &LoginComplete{
			Name: userName,
		}*/

	tpl := &TPL{
		TmpDir: self.TPLDIR,
	}

	res.State = 200
	res.Response = tpl.Login(nil)
}

func (self *Admin) loginFn(req *HTTPServerReq, res *HTTPServerRes) {

}

func (self *Admin) addPost(req *HTTPServerReq, res *HTTPServerRes) {

}

func (self *Admin) notFound(req *HTTPServerReq, res *HTTPServerRes) {
	res.State = 301
	res.Headers = map[string]string{
		"Location": "/",
	}
}
