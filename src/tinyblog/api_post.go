package main

import (
	"strconv"
	"strings"

	"labix.org/v2/mgo/bson"
)

type PostListJson struct {
	S  *Session
	DS *DataService
}

func (self *PostListJson) Get(req *REQ, res *RES) apiResMap {

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

	r.State = true
	r.Data = self.DS.F.TransPostList(pl)
	r.Count = n

	return r.TraceListData()
}

func (self *PostListJson) Set(req *REQ, res *RES) apiResMap {
	r := new(ResJson)
	return r.TraceMsg()
}

func (self *PostListJson) Put(req *REQ, res *RES) apiResMap {
	r := new(ResJson)
	return r.TraceMsg()
}

func (self *PostListJson) Del(req *REQ, res *RES) apiResMap {
	r := new(ResJson)
	return r.TraceMsg()
}

type PostJson struct {
	S  *Session
	DS *DataService
}

func (self *PostJson) Set(req *REQ, res *RES) apiResMap {
	//var rm apiResMap
	r := new(ResJson)
	return r.TraceMsg()
}

func (self *PostJson) Put(req *REQ, res *RES) apiResMap {

	r := new(ResJson)

	uuid := req.GetOneCookieValue("uuid")

	if !self.DS.Auth.HasSavePost(uuid) {
		r.State = false
		r.Message = NOT_ENOUGH_POWER
		return r.TraceMsg()
	}

	title := req.GetFormValue("title")
	if self.DS.Post.IsExist(title) {
		r.State = false
		r.Message = TARGET_HAS_EXIST
		return r.TraceMsg()
	}

	content := req.GetFormValue("content")
	draftVal := req.GetFormValue("draft")
	allowCommentVal := req.GetFormValue("allowcomment")

	tagStr := req.GetFormValue("tags")

	isDraft := false
	if draftVal == "draft" {
		isDraft = true
	}

	allowComment := false
	if allowCommentVal == "allowcomment" {
		allowComment = true
	}

	usr, _ := self.DS.Auth.GetCurUsr(uuid)

	rs := self.DS.Post.Save(&Post{
		Title:        title,
		Content:      content,
		IsDraft:      isDraft,
		AllowComment: allowComment,
		Author:       usr.Name,
		Tags:         strings.Split(tagStr, ","),
	})

	r.State = rs.State
	r.Message = rs.TraceMixMsg()
	return r.TraceMsg()
}

func (self *PostJson) Del(req *REQ, res *RES) apiResMap {

	r := new(ResJson)

	id := req.GetFormValue("id")
	uuid := req.GetOneCookieValue("uuid")

	if self.S.IsLogin(uuid) {

		p, isFound := self.DS.Post.GetOneById(id)

		if isFound {
			if self.DS.Auth.HasEditPost(uuid, p) {
				rs := self.DS.Post.Del(id)
				r.State = rs.State
				r.Message = rs.Message
			} else {
				r.State = false
				r.Message = NOT_ENOUGH_POWER
			}
		} else {
			r.State = false
			r.Message = TARGET_NOT_EXIST
		}

	} else {
		r.State = false
		r.Message = NOT_ENOUGH_POWER
	}

	return r.TraceMsg()
}

func (self *PostJson) Get(req *REQ, res *RES) apiResMap {

	var rm apiResMap
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
			r.Message = NOT_FOUND
			rm = r.TraceMsg()
		}

	} else {
		r.State = false
		r.Message = REQUIRED_DEFAULT
		rm = r.TraceMsg()
	}
	return rm
}
