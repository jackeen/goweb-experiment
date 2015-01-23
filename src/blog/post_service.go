package main

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type PostService struct {
	NumService
	DBC *MDBC
}

func (self *PostService) Insert(p *Post) {

	incId := self.incId(dbc, "post", 1).Post
	currentTime := time.Now()

	data := &Post{
		Id_:          bson.NewObjectId(),
		Id:           incId,
		Title:        p.Title,
		Content:      p.Content,
		Author:       p.Author,
		Cate:         p.Cate,
		Tags:         p.Tags,
		CreateTime:   currentTime,
		LastEditTime: currentTime,
		Draft:        p.Draft,
		AllowComment: p.AllowComment,
		CommentNum:   0,
		CommentIncId: 0,
	}

	self.DBC.Insert(POST_TAB, data)
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

func (self *PostService) InsertComment(postId int, replyId int, content string) {

	post := &Post{}
	self.DBC.SelectOne(POST_TAB, BSONM{"id": postId}, post)
	commentId := post.CommentIncId + 1

	comment := &Comment{
		Id_:        bson.NewObjectId(),
		Id:         commentId,
		Content:    content,
		Auth:       "haha",
		Email:      "e@qq.com",
		host:       "",
		Ip:         "",
		Display:    true,
		ReplyId:    replyId,
		CreateTime: time.Now(),
	}

	self.DBC.UpdatePush(POST_TAB, BSONM{"id": postId}, "comment", comment)
	self.DBC.UpdateInc(POST_TAB, BSONM{"id": postId}, "commentnum", 1)
	self.DBC.UpdateInc(POST_TAB, BSONM{"id": postId}, "commentincid", 1)
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
