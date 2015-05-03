package main

import (
//"log"
//"strconv"
)

type Handler struct {
	Tpl        *TPL
	StaticHost string
	DS         *DataService
	Session    *Session
}

func (self *Handler) Index(req *REQ, res *RES) {

	selData := &SelectData{
		Condition: nil,
		Sort:      "-createtime",
		Limit:     10,
	}

	pl := self.DS.Post.GetList(selData)

	d := map[string]interface{}{
		"PageTitle":  "home",
		"StaticHost": self.StaticHost,
		"PostList":   pl,
	}

	res.Response = self.Tpl.Parse("index", d)
}

func (self *Handler) Post(req *REQ, res *RES) {

	sel := &SelectData{
		Condition: BsonM{
			"title": req.PathParm.FileName,
		},
	}

	p := self.DS.Post.GetOne(sel)

	res.Response = self.Tpl.Parse("post", p)
}

func (self *Handler) Cate(req *REQ, res *RES) {

	res.Response = "post list"
}

func (self *Handler) Tag(req *REQ, res *RES) {

	res.Response = "link"
}

func (self *Handler) Date(req *REQ, res *RES) {

	res.Response = "Date"
}

func (self *Handler) Entry(req *REQ, res *RES) {

	f := req.PathParm.FileName

	switch f {
	case "login":
		self.Login(req, res)
		return
	case "logout":
		self.Logout(req, res)
		return
	}

	d := map[string]interface{}{
		"PageTitle":  "~~login~~",
		"StaticHost": self.StaticHost,
	}

	res.Response = self.Tpl.Parse("login", d)
}

func (self *Handler) Login(req *REQ, res *RES) {

	sel := &SelectData{
		Condition: BsonM{
			"name": req.GetFormValue("user"),
			"pass": req.GetFormValue("pass"),
		},
	}

	user := self.DS.User.GetOne(sel)

	if user.Name == "" {

		GotoLoginErr(req, res)

	} else {

		sd := &SessionData{
			User:  user.Name,
			Power: user.PowerCode,
		}
		uuid := self.Session.New(sd)

		c := res.CreateCookie()
		c.Name = "uuid"
		c.Value = uuid
		c.HttpOnly = true
		c.Path = "/"
		res.SetCookie(c)

		GotoAdminHome(req, res)
	}

}

func (self *Handler) Logout(req *REQ, res *RES) {

}

func (self *Handler) NotFind(req *REQ, res *RES) {

	res.State = 404
	res.Response = "none"
}
