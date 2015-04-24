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

	var postList []Post

	selData := &SelectData{
		Sort:  "-createtime",
		Limit: 10,
		Res:   &postList,
	}

	self.DS.Post.GetList(selData)

	d := map[string]interface{}{
		"PageTitle":  "home",
		"StaticHost": self.StaticHost,
		"PostList":   postList,
	}

	res.Response = self.Tpl.Parse("index", d)
}

func (self *Handler) Post(req *REQ, res *RES) {

	/*var p Post
	sel := BSONM{
		"title": req.PathParm.FileName,
	}

	self.post.SelectOne(sel, &p)

	res.Response = self.Tpl.Parse("post", p)*/

	res.Response = "post info"
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
		Condition{
			"user": req.GetFormValue("user"),
			"pass": req.GetFormValue("pass"),
		},
		Res: &User{},
	}

	self.DS.User.Login(sel)
	//if sel.Err != nil {
	//res.Response = self.Tpl.Parse("loginErr", map[string]string{})
	//}

	user := sel.Res

	if sel.Res.Name == "" {

	} else {

		sd := &SessionData{
			User:  User.Name,
			Power: user.Power,
		}
		uuid = self.session.New(sd)

		c := res.CreateCookie()
		c.Name = "uuid"
		c.Value = uuid
		c.HttpOnly = true
		c.Path = "/"
		res.SetCookie(c)
	}

}

func (self *Handler) Logout(req *REQ, res *RES) {

}

func (self *Handler) NotFind(req *REQ, res *RES) {

	res.State = 404
	res.Response = "none"
}
