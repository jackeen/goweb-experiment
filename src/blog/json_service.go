package main

import (
	"encoding/json"
)

type JsonService struct {
	dbc  *MDBC
	post *PostService
	cate *CateService
}

func (self *JsonService) postByTitle(t string) map[string]interface{} {
	var (
		p    Post
		pMap map[string]interface{}
	)
	sel := Selector{
		"title": t,
	}
	self.post.SelectOne(self.dbc, sel, &p)

	pMap = make(map[string]interface{})
	pMap["title"] = p.Title
	pMap["content"] = p.Content
	pMap["creattime"] = p.CreateTime
	return pMap
}

func (self *JsonService) errorQuery() map[string]interface{} {
	err := map[string]interface{}{
		"success": false,
		"message": "error",
	}
	return err
}

func (self *JsonService) Init(dbc *MDBC) {
	self.dbc = dbc
}

func (self *JsonService) GetJson(req *HTTPServerReq, res *HTTPServerRes) {
	//m := req.Query["m"]
	v, _ := json.Marshal(self.errorQuery())
	res.State = 404
	res.Response = string(v)
}
