package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type PostService struct {
	DBC *MDBC
	C   *mgo.Collection
	S   *Session
}

func (self *PostService) Save(p *Post) *ResMessage {

	currentTime := time.Now()

	if p.Title == "" || p.Content == "" {
		return getUserResMessage(false, REQUIRED_DEFAULT, POST_MODE_CODE)
	}

	p.Id_ = bson.NewObjectId()
	p.CreateTime = currentTime
	p.EditTime = currentTime

	err := self.C.Insert(p)

	return getResMessage(err, SAVE_SUCCESS, POST_MODE_CODE)
}

func (self *PostService) GetOne(sel *SelectData) *Post {

	p := new(Post)
	self.C.Find(sel.Condition).One(p)
	return p
}

func (self *PostService) GetList(sel *SelectData) *PostList {

	pl := &PostList{}
	q := self.C.Find(sel.Condition)
	q = q.Sort(sel.Sort).Limit(sel.Limit)
	q.All(pl)
	return pl
}

/*


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
*/
//func (self *PostService) updateComment( postId int, commentId int, BSONM sel) {

//db.shcool.update({ "_id" : 2, "students.name" : "ajax"},{"$inc" : {"students.0.age" : 1} });
//postSel := BSONM{"id": postId}
//commentSel := BSONM{"id": commentId}
//self.DBC.UpdateSet(POST_TAB, postSel, data)

//}
