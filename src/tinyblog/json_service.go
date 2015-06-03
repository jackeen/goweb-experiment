package main

import (
	"encoding/json"
	"labix.org/v2/mgo/bson"
	//"log"
	"strconv"
	//"strings"
)

type ResJsonMap map[string]interface{}

type ResJson struct {
	State bool
	Msg   string
	Count int
	Data  interface{}
}

func (self *ResJson) TraceMsg() ResJsonMap {
	return ResJsonMap{
		"state": self.State,
		"msg":   self.Msg,
	}
}

func (self *ResJson) TraceNotFound() ResJsonMap {
	self.State = false
	self.Msg = NOT_FOUND
	return self.TraceMsg()
}

func (self *ResJson) TraceData() ResJsonMap {
	return ResJsonMap{
		"state": self.State,
		"data":  self.Data,
	}
}

func (self *ResJson) TraceListData() ResJsonMap {
	return ResJsonMap{
		"state": self.State,
		"count": self.Count,
		"data":  self.Data,
	}
}

type IJson interface {
	Get(*REQ, *RES) ResJsonMap
	Set(*REQ, *RES) ResJsonMap
	Put(*REQ, *RES) ResJsonMap
	Del(*REQ, *RES) ResJsonMap
}

//cate api
type CateJson struct {
	S  *Session
	DS *DataService
}

func (self *CateJson) Get(req *REQ, res *RES) ResJsonMap {

	r := new(ResJson)
	qParent := req.GetFormValue("p")

	cs := self.DS.Cate.GetNames(qParent)
	r.State = true
	r.Data = cs
	r.Count = len(cs)
	return r.TraceListData()
}

func (self *CateJson) Set(req *REQ, res *RES) ResJsonMap {
	r := new(ResJson)
	return r.TraceMsg()
}

func (self *CateJson) Put(req *REQ, res *RES) ResJsonMap {

	r := new(ResJson)
	qName := req.GetFormValue("n")
	qParent := req.GetFormValue("p")

	if qName == "" {
		r.State = false
		r.Msg = REQUIRED_DEFAULT
		return r.TraceMsg()
	}

	c := &Cate{
		Name:   qName,
		Parent: qParent,
	}
	rs := self.DS.Cate.Save(c)

	r.State = rs.State
	r.Msg = rs.TraceMixMsg()
	return r.TraceMsg()
}

func (self *CateJson) Del(req *REQ, res *RES) ResJsonMap {
	r := new(ResJson)
	return r.TraceMsg()
}

//user api
type UserJson struct {
	S  *Session
	DS *DataService
}

func (self *UserJson) Get(req *REQ, res *RES) ResJsonMap {
	r := new(ResJson)
	return r.TraceMsg()
}

func (self *UserJson) Set(req *REQ, res *RES) ResJsonMap {
	r := new(ResJson)
	return r.TraceMsg()
}

func (self *UserJson) Put(req *REQ, res *RES) ResJsonMap {
	r := new(ResJson)
	return r.TraceMsg()
}

func (self *UserJson) Del(req *REQ, res *RES) ResJsonMap {
	r := new(ResJson)
	return r.TraceMsg()
}

type PostListJson struct {
	S  *Session
	DS *DataService
}

func (self *PostListJson) Get(req *REQ, res *RES) ResJsonMap {

	r := new(ResJson)
	page := 0
	limit := 5

	reqTitle := req.GetFormValue("t")

	reqPage := req.GetFormValue("p")
	reqLimit := req.GetFormValue("l")

	if reqPage != "" {
		p, err := strconv.ParseInt(reqPage, 10, 32)
		if err == nil {
			page = int(p)
		}
	}

	if reqLimit != "" {
		l, err := strconv.ParseInt(reqLimit, 10, 32)
		if err == nil && l < 5 {
			limit = int(l)
		}
	}

	selData := &SelectData{
		Condition: nil,
		Sort:      "-createtime",
		Start:     page * limit,
		Limit:     limit,
	}

	if reqTitle != "" {
		selData.Condition = bson.M{"title": bson.M{"$regex": bson.RegEx{reqTitle, "i"}}}
	}

	pl := self.DS.Post.GetList(selData)
	n := self.DS.Post.Count(selData)

	f := new(Format)
	pLen := len(pl)
	plm := make([]map[string]interface{}, 0)
	for i := 0; i < pLen; i++ {
		plm = append(plm, f.O2M(pl[i]))
	}

	r.State = true
	r.Data = plm
	r.Count = n

	return r.TraceListData()
}

func (self *PostListJson) Set(req *REQ, res *RES) ResJsonMap {
	r := new(ResJson)
	return r.TraceMsg()
}

func (self *PostListJson) Put(req *REQ, res *RES) ResJsonMap {
	r := new(ResJson)
	return r.TraceMsg()
}

func (self *PostListJson) Del(req *REQ, res *RES) ResJsonMap {
	r := new(ResJson)
	return r.TraceMsg()
}

type PostJson struct {
	S  *Session
	DS *DataService
}

func (self *PostJson) Set(req *REQ, res *RES) ResJsonMap {
	//var rm ResJsonMap
	r := new(ResJson)
	return r.TraceMsg()
}

func (self *PostJson) Put(req *REQ, res *RES) ResJsonMap {

	r := new(ResJson)

	title := req.GetFormValue("title")
	content := req.GetFormValue("content")
	draftVal := req.GetFormValue("draft")
	allowCommentVal := req.GetFormValue("allowcomment")

	isDraft := false
	if draftVal == "draft" {
		isDraft = true
	}

	allowComment := false
	if allowCommentVal == "allowcomment" {
		allowComment = true
	}

	rs := self.DS.Post.Save(&Post{
		Title:        title,
		Content:      content,
		IsDraft:      isDraft,
		AllowComment: allowComment,
	})

	r.State = rs.State
	r.Msg = rs.TraceMixMsg()
	return r.TraceMsg()
}

func (self *PostJson) Del(req *REQ, res *RES) ResJsonMap {
	//var rm ResJsonMap
	r := new(ResJson)

	//uuid := self.S.Get(req.GetCookieValues("uuid"))

	return r.TraceMsg()
}

func (self *PostJson) Get(req *REQ, res *RES) ResJsonMap {

	var rm ResJsonMap
	r := new(ResJson)
	t := req.GetUrlOneValue("t")

	if t != "" {

		p := self.DS.Post.GetOne(&SelectData{
			Condition: bson.M{
				"title": t,
			},
		})

		if p.Title != "" {
			f := new(Format)
			r.Data = f.O2M(*p)
			r.State = true
			rm = r.TraceData()
		} else {
			r.State = false
			r.Msg = NOT_FOUND
			rm = r.TraceMsg()
		}

	} else {
		r.State = false
		r.Msg = REQUIRED_DEFAULT
		rm = r.TraceMsg()
	}
	return rm
}

type JsonService struct {
	S  *Session
	DS *DataService
}

func (self *JsonService) matchFn(obj IJson, req *REQ, res *RES) ResJsonMap {
	var resJson ResJsonMap
	switch req.PathParm.FileName {
	case "get":
		resJson = obj.Get(req, res)
	case "put":
		resJson = obj.Put(req, res)
	default:
		resJson = new(ResJson).TraceNotFound()
	}
	return resJson
}

func (self *JsonService) Cate(req *REQ, res *RES) ResJsonMap {
	return self.matchFn(&CateJson{
		S:  self.S,
		DS: self.DS,
	}, req, res)
}

func (self *JsonService) User(req *REQ, res *RES) ResJsonMap {
	return self.matchFn(&UserJson{
		S:  self.S,
		DS: self.DS,
	}, req, res)
}

func (self *JsonService) PostList(req *REQ, res *RES) ResJsonMap {
	return self.matchFn(&PostListJson{
		S:  self.S,
		DS: self.DS,
	}, req, res)
}

func (self *JsonService) Post(req *REQ, res *RES) ResJsonMap {
	return self.matchFn(&PostJson{
		S:  self.S,
		DS: self.DS,
	}, req, res)
}

func (self *JsonService) Rout(req *REQ, res *RES) {

	var resJson ResJsonMap

	p := req.PathParm.PathItems

	if len(p) == 2 {
		switch p[1] {
		case "post":
			resJson = self.Post(req, res)
		case "postlist":
			resJson = self.PostList(req, res)
		case "user":
			resJson = self.User(req, res)
		case "cate":
			resJson = self.Cate(req, res)
		default:
			resJson = new(ResJson).TraceNotFound()
		}
	} else {
		resJson = new(ResJson).TraceNotFound()
	}

	v, _ := json.Marshal(resJson)
	res.SetHeader("Content-Type", "application/json;charset=UTF-8")
	res.Response = string(v)
}
