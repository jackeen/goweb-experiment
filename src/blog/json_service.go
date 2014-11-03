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

func (self *JsonService) Init(dbc *MDBC) {
	self.dbc = dbc
	self.postService = &PostService{}
	self.cateService = &CateService{}
}

func (self *JsonService) postInfo(t []string) map[string]interface{} {
	var (
		p       Post
		jsonMap map[string]interface{}
	)

	log.Println("xxxx", self.postService)
	log.Println(t)

	sel := Selector{
		"title": t[0],
	}

	self.postService.SelectOne(self.dbc, sel, &p)

	log.Println(p)

	jsonMap = map[string]interface{}{
		"title":      p.Title,
		"content":    p.Content,
		"createtime": p.CreateTime,
	}

	/*if len(t) == 1 {

		sel := Selector{
			"title": t[0],
		}

		log.Println(sel)

		self.postService.SelectOne(self.dbc, sel, &p)

		jsonMap = map[string]interface{}{
			"title":      p.Title,
			"content":    p.Content,
			"createtime": p.CreateTime,
		}
	} else {
		jsonMap = self.errorQuery()
	}*/

	return jsonMap
}

func (self *JsonService) errorQuery() map[string]interface{} {
	err := map[string]interface{}{
		"success": false,
		"message": "not found",
	}
	return err
}

func (self *JsonService) GetJson(req *HTTPServerReq, res *HTTPServerRes) {

	var (
		queryJson map[string]interface{}
	)

	switch req.PathParm.FileName {
	case "post":
		queryJson = self.postInfo(req.Query["title"])
	default:
		queryJson = self.errorQuery()
	}

	v, _ := json.Marshal(queryJson)
	res.State = 200
	res.Response = string(v)
}
