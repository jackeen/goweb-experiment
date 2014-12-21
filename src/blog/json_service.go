package main

import (
	"encoding/json"
	//"log"
)

const (
	DateFormatStr = "2006-01-02 15:04:05"
)

type JsonService struct {
	session     *Session
	dbc         *MDBC
	postService *PostService
	cateService *CateService
}

func (self *JsonService) Init(dbc *MDBC, s *Session) {
	self.session = s
	self.dbc = dbc
	self.postService = &PostService{}
	self.cateService = &CateService{}
}

func (self *JsonService) postInfo(t string) map[string]interface{} {
	var (
		p       Post
		jsonMap map[string]interface{}
	)

	if t != "" {

		sel := Selector{
			"title": t,
		}

		self.postService.SelectOne(self.dbc, sel, &p)

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

func (self *JsonService) login(req *REQ, res *RES) map[string]interface{} {

	u := req.GetFormValue("user")
	p := req.GetFormValue("pass")

	user := &User{}
	uuid := req.GetCookies()["uuid"]

	m := make(map[string]interface{})

	if self.session.Get(uuid) != nil {
		m["success"] = true
		m["message"] = "ready!"
		return m
	}

	uc := &UserService{}
	uc.LoginSelect(self.dbc, u, p, user)

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
	case "post":
		queryJson = self.postInfo(req.GetUrlOneValue("t"))
	case "login":
		queryJson = self.login(req, res)
	default:
		queryJson = self.errorQuery()
	}

	v, _ := json.Marshal(queryJson)
	res.State = 200
	res.Response = string(v)
}
