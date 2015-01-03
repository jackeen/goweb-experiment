package main

import (
	//"log"
	//"reflect"
	//"errors"
	//"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

const (
	NUM_TAB    = "num"
	POST_TAB   = "post"
	CATE_TAB   = "cate"
	USER_TAB   = "user"
	TAG_TAB    = "tag"
	NAV_TAB    = "nav"
	CONFIG_TAB = "config"
)

//inc id num data I/O
type NumService struct{}

func (self *NumService) Init(dbc *MDBC) {
	dbc.Insert(NUM_TAB, &Num{-1, -1, -1, -1})
}

func (self *NumService) incId(dbc *MDBC, colName string, i int) *Num {
	res := &Num{}
	dbc.UpdateInc(NUM_TAB, nil, colName, i)
	dbc.SelectOne(NUM_TAB, nil, res)
	return res
}

var IncNum *NumService = new(NumService)

//user data IO
type UserService struct{}

func (self *UserService) Insert(dbc *MDBC, name string, pass string, nick string, email string, power int) {

	user := &User{
		Id_:        bson.NewObjectId(),
		Id:         IncNum.incId(dbc, "user", 1).User,
		Name:       name,
		Pass:       pass,
		Nick:       nick,
		Email:      email,
		Power:      power,
		CreateTime: time.Now(),
	}
	dbc.Insert(USER_TAB, user)
}

func (self *UserService) Update(dbc *MDBC, sel Selector, data interface{}) {
	dbc.UpdateSet(USER_TAB, sel, data)
}

func (self *UserService) Select(dbc *MDBC) {

}

func (self *UserService) Delete(dbc *MDBC) {

}

func (self *UserService) LoginSelect(dbc *MDBC, u string, p string, res *User) {
	sel := Selector{
		"name": u,
		"pass": p,
	}
	dbc.SelectOne(USER_TAB, sel, res)
}

func (self *UserService) HasUser(dbc *MDBC, name string) bool {
	res := &User{}
	dbc.SelectOne(USER_TAB, Selector{"name": name}, res)
	if res.Name == "" {
		return false
	} else {
		return true
	}
}

//
type CateService struct{}

func (self *CateService) Inert(dbc *MDBC, name string, exp string, pid int) {

	cate := &Cate{
		Id_:      bson.NewObjectId(),
		Id:       IncNum.incId(dbc, "cate", 1).Cate,
		Name:     name,
		Explain:  exp,
		ParentId: pid,
	}
	dbc.Insert(CATE_TAB, cate)
}

func (self *CateService) Select(dbc *MDBC, id int, cate *Cate) {

	dbc.SelectOne(CATE_TAB, Selector{"id": id}, cate)
}

func (self *CateService) Update(dbc *MDBC, id int, data interface{}) {

	dbc.UpdateSet(CATE_TAB, Selector{"id": id}, data)
}

func (self *CateService) Delete(dbc *MDBC, id int) {
	dbc.Delete(CATE_TAB, Selector{"id": id})
}

//post
type PostService struct{}

func (self *PostService) Insert(dbc *MDBC, p *Post) {

	incId := IncNum.incId(dbc, "post", 1).Post
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
		EditState:    p.EditState,
		AllowComment: p.AllowComment,
		CommentNum:   0,
		CommentIncId: 0,
	}
	dbc.Insert(POST_TAB, data)
}

func (self *PostService) Select(dbc *MDBC,
	sel Selector,
	sort string,
	offset int,
	limit int,
	res *[]Post) {

	dbc.Select(POST_TAB, sel, sort, offset, limit, res)
}

func (self *PostService) SelectOne(dbc *MDBC, sel Selector, res *Post) {
	dbc.SelectOne(POST_TAB, sel, res)
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

//func (self *PostService) updateComment(dbc *MDBC, postId int, commentId int, Selector sel) {

//db.shcool.update({ "_id" : 2, "students.name" : "ajax"},{"$inc" : {"students.0.age" : 1} });
//postSel := Selector{"id": postId}
//commentSel := Selector{"id": commentId}
//dbc.UpdateSet(POST_TAB, postSel, data)

//}
