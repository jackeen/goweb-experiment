package main

import (
	//"log"
	//"strconv"
	"net/http"
	//"time"
	"labix.org/v2/mgo/bson"
)

type PageData struct {
	Title      string
	StaticHost string
	Data       interface{}
	Sub        string
}

type Handler struct {
	Tpl        *TPL
	StaticHost string
	DS         *DataService
	Session    *Session
}

func (self *Handler) GetPD(t string, d interface{}) PageData {
	return PageData{
		Title:      t,
		StaticHost: self.StaticHost,
		Data:       d,
		Sub:        "footer",
	}
}

func (self *Handler) Index(req *REQ, res *RES) {

	selData := &SelectData{
		Condition: nil,
		Sort:      "-createtime",
		Limit:     10,
	}

	pl := self.DS.Post.GetList(selData)

	d := self.GetPD("tiny", self.DS.F.TranPost(pl))
	res.Response = self.Tpl.Parse("index", d)
}

func (self *Handler) Post(req *REQ, res *RES) {

	fileName := req.PathParm.FileName

	sel := &SelectData{}

	if bson.IsObjectIdHex(fileName) {
		sel.Condition = bson.M{
			"_id": bson.ObjectIdHex(fileName),
		}
	} else {
		sel.Condition = bson.M{
			"title": req.PathParm.FileName,
		}
	}

	p := self.DS.Post.GetOne(sel)

	if p.Title == "" {
		//404
	}

	d := self.GetPD(p.Title, self.DS.F.O2M(*p))
	res.Response = self.Tpl.Parse("post", d)
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

	d := self.GetPD("login", nil)

	res.Response = self.Tpl.Parse("login", d)
}

func (self *Handler) Login(req *REQ, res *RES) {

	sel := &SelectData{
		Condition: bson.M{
			"name": req.GetFormValue("user"),
			"pass": req.GetFormValue("pass"),
		},
	}

	user := self.DS.User.GetOne(sel)

	if user.Name == "" {

		GotoLoginErr(req, res)

	} else {

		sd := &SessionData{
			U: user,
		}
		uuid := self.Session.New(sd)

		c := &http.Cookie{
			Name:     "uuid",
			Value:    uuid,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(res.W, c)
		GotoAdminHome(req, res)
	}

}

func (self *Handler) Logout(req *REQ, res *RES) {
	c, err := req.R.Cookie("uuid")
	if err == nil {
		self.Session.Destroy(c.Value)
	}

	GotoHome(req, res)
}

func (self *Handler) NotFind(req *REQ, res *RES) {

	res.State = 404
	res.Response = "none"
}
