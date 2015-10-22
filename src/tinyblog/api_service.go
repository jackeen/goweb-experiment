package main

import (
	"encoding/json"
	//"labix.org/v2/mgo/bson"
	//"log"
	//"ioutil"

	//"strconv"
	//"strings"
)

type apiResMap map[string]interface{}

type ResJson struct {
	State   bool
	Message string
	Count   int
	Data    interface{}
}

func (self *ResJson) TraceMsg() apiResMap {
	return apiResMap{
		"state":   self.State,
		"message": self.Message,
	}
}

func (self *ResJson) TraceNotFound() apiResMap {
	self.State = false
	self.Message = NOT_FOUND
	return self.TraceMsg()
}

func (self *ResJson) TraceData() apiResMap {
	return apiResMap{
		"state": self.State,
		"data":  self.Data,
	}
}

func (self *ResJson) TraceListData() apiResMap {
	return apiResMap{
		"state": self.State,
		"count": self.Count,
		"data":  self.Data,
	}
}

type IJson interface {
	Get(*REQ, *RES) apiResMap
	Set(*REQ, *RES) apiResMap
	Put(*REQ, *RES) apiResMap
	Del(*REQ, *RES) apiResMap
}

/*public router*/
type APIService struct {
	S  *Session
	DS *DataService
}

func (self *APIService) matchFn(obj IJson, req *REQ, res *RES) apiResMap {
	var resJson apiResMap
	switch req.PathParm.FileName {
	case "get":
		resJson = obj.Get(req, res)
	case "put":
		resJson = obj.Put(req, res)
	case "del":
		resJson = obj.Del(req, res)
	default:
		resJson = new(ResJson).TraceNotFound()
	}
	return resJson
}

func (self *APIService) Tag(req *REQ, res *RES) apiResMap {
	return self.matchFn(&TagJson{
		S:  self.S,
		DS: self.DS,
	}, req, res)
}

func (self *APIService) Cate(req *REQ, res *RES) apiResMap {
	return self.matchFn(&CateJson{
		S:  self.S,
		DS: self.DS,
	}, req, res)
}

func (self *APIService) User(req *REQ, res *RES) apiResMap {
	return self.matchFn(&UserJson{
		S:  self.S,
		DS: self.DS,
	}, req, res)
}

func (self *APIService) PostList(req *REQ, res *RES) apiResMap {
	return self.matchFn(&PostListJson{
		S:  self.S,
		DS: self.DS,
	}, req, res)
}

func (self *APIService) Post(req *REQ, res *RES) apiResMap {
	return self.matchFn(&PostJson{
		S:  self.S,
		DS: self.DS,
	}, req, res)
}

func (self *APIService) Image(req *REQ, res *RES) apiResMap {
	return self.matchFn(&ImageAPI{
		S:  self.S,
		DS: self.DS,
	}, req, res)
}

func (self *APIService) Rout(req *REQ, res *RES) {

	var resJson apiResMap

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
		case "tag":
			resJson = self.Tag(req, res)
		case "image":
			resJson = self.Image(req, res)
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
