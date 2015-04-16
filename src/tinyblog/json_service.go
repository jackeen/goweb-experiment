package main

import (
	"encoding/json"
	//"log"
)

type JsonService struct {
	Session *Session
	DS      *DataService
}

func (self *JsonService) getPost(req *REQ, res *RES) map[string]interface{} {
	var (
		p       Post
		jsonMap map[string]interface{}
	)

	t := req.GetUrlOneValue("t")

	if t != "" {

		sel := BSONM{
			"title": t,
		}

		self.postService.SelectOne(sel, &p)

		jsonMap = map[string]interface{}{
			"title":      p.Title,
			"content":    p.Content,
			"createtime": p.CreateTime.Format(DateFormatStr),
		}
	} else {
		jsonMap = self.errorQuery()
	}

	return jsonMap
}

func (self *JsonService) savePost(req *REQ, res *RES) ResMessage {

	title := req.GetFormValue("title")
	content := req.GetFormValue("content")
	draft := req.GetFormValue("draft")

	isDraft := false
	if draft == "draft" {
		isDraft = true
	}

	p := new(Post)
	p.Title = title
	p.Content = content
	p.Draft = isDraft
	rs := self.postService.Insert(p)

	return rs
}

func (self *JsonService) login(req *REQ, res *RES) map[string]interface{} {

	u := req.GetFormValue("user")
	p := req.GetFormValue("pass")

	user := &User{}
	uuid := req.GetCookies()["uuid"]

	m := make(map[string]interface{})

	if self.session.IsLogin(uuid) {
		m["success"] = true
		m["message"] = "ready!"
		return m
	}

	uc := &UserService{}
	uc.LoginSelect(u, p, user)

	if user.Name == "" {
		m["success"] = false
		m["message"] = "user or pass error"
	} else {
		m["success"] = true
		m["message"] = "welcome"

		sd := &SessionData{
			User:  user.Name,
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
	return m
}

func (self *JsonService) logout() {

}

func (self *JsonService) errorQuery() map[string]interface{} {
	err := map[string]interface{}{
		"success": false,
		"message": "not found",
	}
	return err
}

func (self *JsonService) GetJson(req *REQ, res *RES) {

	var (
		queryJson map[string]interface{}
	)

	switch req.PathParm.FileName {
	case "getpost":
		queryJson = self.getPost(req, res)
	case "savepost":
		queryJson = self.savePost(req, res)
	case "login":
		queryJson = self.login(req, res)
	default:
		queryJson = self.errorQuery()
	}

	v, _ := json.Marshal(queryJson)
	res.Response = string(v)
}
