package main

import (
	"encoding/json"
	//"reflect"
	//"log"
)

type ResJsonMap map[string]interface{}

type ResJson struct {
	State bool
	Msg   string
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

type IJson interface {
	Get(*REQ, *RES) ResJsonMap
	Set(*REQ, *RES) ResJsonMap
	Put(*REQ, *RES) ResJsonMap
	Del(*REQ, *RES) ResJsonMap
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
	draft := req.GetFormValue("draft")

	isDraft := false
	if draft == "draft" {
		isDraft = true
	}

	rs := self.DS.Post.Save(&Post{
		Title:   title,
		Content: content,
		IsDraft: isDraft,
	})

	r.State = rs.State
	r.Msg = rs.TraceMixMsg()

	return r.TraceMsg()
}

func (self *PostJson) Del(req *REQ, res *RES) ResJsonMap {
	//var rm ResJsonMap
	r := new(ResJson)
	return r.TraceMsg()
}

func (self *PostJson) Get(req *REQ, res *RES) ResJsonMap {

	var rm ResJsonMap
	r := new(ResJson)
	t := req.GetUrlOneValue("t")

	if t != "" {

		p := self.DS.Post.GetOne(&SelectData{
			Condition: BsonM{
				"title": t,
			},
		})

		if p.Title != "" {
			d := ResJsonMap{
				"title":      p.Title,
				"content":    p.Content,
				"createtime": p.CreateTime.Format(DateFormatStr),
				"isdraft":    p.IsDraft,
			}
			r.Data = d
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
		default:
			resJson = new(ResJson).TraceNotFound()
		}
	} else {
		resJson = new(ResJson).TraceNotFound()
	}

	v, _ := json.Marshal(resJson)
	res.Response = string(v)

}
