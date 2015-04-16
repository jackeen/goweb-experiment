package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type PostService struct {
	DBC *MDBC
	C   *mgo.Collection
}

func (self *PostService) Insert(p *Post) ResMessage {

	currentTime := time.Now()

	rs := new(ResMessage)

	if p.Title == "" || p.Content == "" {
		rs.State = false
		rs.Message = SaveDataFail
		return rs
	}

	data := &Post{
		Id_:          bson.NewObjectId(),
		Title:        p.Title,
		Content:      p.Content,
		Author:       p.Author,
		Cate:         p.Cate,
		Tags:         p.Tags,
		CreateTime:   currentTime,
		EditTime:     currentTime,
		IsDraft:      p.Draft,
		AllowComment: p.AllowComment,
	}

	err := self.DBC.Insert(POST_TAB, data)

	if err == nil {
		rs.State = true
		rs.Message = SaveSuccess
	} else {
		rs.State = false
		rs.Message = err.Error()
	}

	return rs

}

func (self *PostService) Select(sel *SelectData) {
	sel.Tab = POST_TAB
	self.DBC.Select(sel)
}

func (self *PostService) SelectOne(sel BSONM, res *Post) {
	self.DBC.SelectOne(POST_TAB, sel, res)
}

func (self *PostService) Update(postId int, data interface{}) {
	self.DBC.UpdateSet(POST_TAB, BSONM{"id": postId}, data)
}

func (self *PostService) InsertComment(postId int, content string) {

	post := &Post{}
	self.DBC.SelectOne(POST_TAB, BSONM{"id": postId}, post)

	comment := &Comment{
		Content:    content,
		Auth:       "haha",
		Email:      "e@qq.com",
		host:       "",
		Ip:         "",
		Display:    true,
		CreateTime: time.Now(),
	}

	self.DBC.UpdatePush(POST_TAB, BSONM{"id": postId}, "comment", comment)

}

func (self *PostService) deleteComment(postId int, commentId int) {

	postSel := BSONM{"id": postId}
	commentSel := BSONM{"id": commentId}
	self.DBC.UpdatePull(POST_TAB, postSel, "comment", commentSel)
	self.DBC.UpdateInc(POST_TAB, BSONM{"id": postId}, "commentnum", -1)
}

//func (self *PostService) updateComment( postId int, commentId int, BSONM sel) {

//db.shcool.update({ "_id" : 2, "students.name" : "ajax"},{"$inc" : {"students.0.age" : 1} });
//postSel := BSONM{"id": postId}
//commentSel := BSONM{"id": commentId}
//self.DBC.UpdateSet(POST_TAB, postSel, data)

//}
