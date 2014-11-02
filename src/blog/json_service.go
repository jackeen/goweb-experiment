package main

import (
	"encoding/json"
	"log"
)

type JsonService struct {
	dbc         *MDBC
	postService *PostService
	cateService *CateService
}

func (self *JsonService) post(q map[string][]string) map[string]interface{} {
	var (
		title   string
		p       Post
		jsonMap map[string]interface{}
	)

	t := q["title"]

	log.Println(t)

	if len(t) == 1 {
		title = t[0]

		log.Println(title)

		self.postService.SelectOne(self.dbc, Selector{"title": title}, &p)

		log.Println(p)

		jsonMap = map[string]interface{}{
			"title":      p.Title,
			"content":    p.Content,
			"createtime": p.CreateTime,
		}
	} else {
		jsonMap = self.errorQuery()
	}

	return jsonMap
}

func (self *JsonService) errorQuery() map[string]interface{} {
	err := map[string]interface{}{
		"success": false,
		"message": "not found",
	}
	return err
}

func (self *JsonService) Init(dbc *MDBC) {
	self.dbc = dbc
	self.postService = &PostService{}
	self.cateService = &CateService{}
}

func (self *JsonService) GetJson(req *HTTPServerReq, res *HTTPServerRes) {

	var (
		queryJson map[string]interface{}
	)

	switch req.PathParm.FileName {
	case "post":
		queryJson = self.post(req.Query)
	default:
		queryJson = self.errorQuery()
	}

	v, _ := json.Marshal(queryJson)
	res.State = 200
	res.Response = string(v)
}
