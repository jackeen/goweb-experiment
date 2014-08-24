package main

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type PostService struct{}

func (self *PostService) Insert(dbc *MDBC, title string, content string, author int, cate int, tags []string, isEdit bool, allowComment bool) {

	incId := IncNum.incId(dbc, "post", 1).Post
	currentTime := time.Now()

	data := &Post{
		Id_:          bson.NewObjectId(),
		Id:           incId,
		Title:        title,
		Content:      content,
		Author:       author,
		Cate:         cate,
		Tags:         tags,
		CreateTime:   currentTime,
		LastEditTime: currentTime,
		EditState:    isEdit,
		AllowComment: allowComment,
		CommentNum:   0,
		CommentIncId: 0,
	}
	dbc.Insert(POST_TAB, data)
}

func (self *PostService) Select(dbc *MDBC, sel Selector, sort string, offset int, limit int, res *[]Post) {

	dbc.Select(POST_TAB, sel, sort, offset, limit, res)
}

func (self *PostService) Update(dbc *MDBC, postId int, data interface{}) {
	dbc.UpdateSet(POST_TAB, Selector{"id": postId}, data)
}

func (self *PostService) InsertComment(dbc *MDBC, postId int, replyId int, content string) {

	post := &Post{}
	dbc.SelectOne(POST_TAB, Selector{"id": postId}, post)
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

	dbc.UpdatePush(POST_TAB, Selector{"id": postId}, "comment", comment)

	dbc.UpdateInc(POST_TAB, Selector{"id": postId}, "commentnum", 1)
	dbc.UpdateInc(POST_TAB, Selector{"id": postId}, "commentincid", 1)
}

func (self *PostService) deleteComment(dbc *MDBC, postId int, commentId int) {

	postSel := Selector{"id": postId}
	commentSel := Selector{"id": commentId}
	dbc.UpdatePull(POST_TAB, postSel, "comment", commentSel)
	dbc.UpdateInc(POST_TAB, Selector{"id": postId}, "commentnum", -1)
}

func (self *PostService) updateComment(dbc *MDBC, postId int, commentId int, Selector sel) {
	postSel := Selector{"id": postId}
	commentSel := Selector{"id": commentId}

}
