package main

import (
	"encoding/json"

	//"log"
)

type jsonMap map[string]interface{}

type JsonService struct {
	Session *Session
	DS      *DataService
}

func (self *JsonService) getPost(req *REQ, res *RES) jsonMap {
	var (
		p *Post
		m jsonMap
	)

	t := req.GetUrlOneValue("t")

	if t != "" {

		p = &Post{}
		self.DS.Post.GetOneByTitle(t, p)

		m = jsonMap{
			"title":      p.Title,
			"content":    p.Content,
			"createtime": p.CreateTime.Format(DateFormatStr),
		}
	} else {
		m = self.errorQuery()
	}

	return m
}

func (self *JsonService) savePost(req *REQ, res *RES) jsonMap {

	title := req.GetFormValue("title")
	content := req.GetFormValue("content")
	draft := req.GetFormValue("draft")

	isDraft := false
	if draft == "draft" {
		isDraft = true
	}

	rs := self.DS.Post.Insert(&Post{
		Title:   title,
		Content: content,
		IsDraft: isDraft,
	})

	return jsonMap{
		"state":   rs.State,
		"message": rs.Message,
	}
}

/*
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
}*/

func (self *JsonService) logout() {

}

func (self *JsonService) errorQuery() jsonMap {
	err := jsonMap{
		"success": false,
		"message": "not found",
	}
	return err
}

func (self *JsonService) GetJson(req *REQ, res *RES) {

	var (
		queryJson jsonMap
	)

	switch req.PathParm.FileName {
	case "getpost":
		queryJson = self.getPost(req, res)
	case "savepost":
		queryJson = self.savePost(req, res)
	//case "login":
	//	queryJson = self.login(req, res)
	default:
		queryJson = self.errorQuery()
	}

	v, _ := json.Marshal(queryJson)
	res.Response = string(v)

}
