package main

import (
	"encoding/json"
	"log"
)

type Admin struct {
	DBC        *MDBC
	TPL        *AdminTPL
	StaticHost string
	session    *Session
	entry      *EntryPage
}

func (self *Admin) Init() {

	self.session = &Session{
		Data: make(map[string]*SessionData),
	}

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

func (self *Admin) Auth(u string, p string) int {

}

func (self *Admin) AdminNotFoundPage(req *REQ, res *RES) {
	res.State = 404
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
	d := &EntryPageData{
		StaticHost: self.Parent.StaticHost,
	}
	log.Println(d)
	res.State = 200
	res.Response = self.Parent.TPL.Login(d)
}

func (self *EntryPage) loginServe(req *REQ, res *RES) {

	u := req.GetFormValue("user")
	p := req.GetFormValue("pass")

	user := &User{}
	uuid := req.GetCookies()["uuid"]

	m := make(map[string]interface{})

	if self.Parent.session.Get(uuid) != nil {
		m["success"] = true
		m["message"] = "ready!"
		b, _ := json.Marshal(m)
		res.State = 200
		res.Response = string(b)
		return
	}

	uc := &UserService{}
	uc.LoginSelect(self.Parent.DBC, u, p, user)

	if user.Name == "" {
		m["success"] = false
		m["message"] = "user or pass error"
	} else {
		m["success"] = true
		m["message"] = "welcome"

		sd := &SessionData{
			User: user.Name,
		}
		uuid = self.Parent.session.New(sd)

		c := res.CreateCookie()
		c.Name = "uuid"
		c.Value = uuid
		c.HttpOnly = true
		c.Path = "/"
		res.SetCookie(c)
	}
	b, _ := json.Marshal(m)

	res.State = 200
	res.Response = string(b)
}

func (self *EntryPage) logoutServe(req *REQ, res *RES) {

}

type EntryPageData struct {
	StaticHost string
}
