package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	//"time"
)

const (
	POST_TAB   = "post"
	CATE_TAB   = "cate"
	USER_TAB   = "user"
	TAG_TAB    = "tag"
	CONFIG_TAB = "config"
)

const (
	SAVE_SUCCESS     = "save success"
	SAVE_FAIL        = "save fail"
	REQUIRED_DEFAULT = "required default"
)

const (
	POST_MODE_CODE = 101
	USER_MODE_CODE = 102
)

type ResMessage struct {
	State   bool
	Addr    int
	Message string
}

type BsonM bson.M

type SelectData struct {
	Condition BsonM
	Sort      string
	Limit     int
	GT        string
}

func getResMessage(err error, msg string, n int) *ResMessage {

	rs := new(ResMessage)
	if err == nil {
		rs.State = true
		rs.Message = msg
	} else {
		rs.State = false
		rs.Message = err.Error()
	}
	rs.Addr = n
	return rs
}

func getUserResMessage(s bool, msg string, n int) *ResMessage {
	return &ResMessage{
		State:   s,
		Message: msg,
		Addr:    n,
	}
}

//mdb connection
type MDBC struct {
	Host string
	User string
	Pass string
	Name string
	S    *mgo.Session
	DB   *mgo.Database
}

func (self *MDBC) Init() {

	dbQuery := "mongodb://" + self.User + ":" + self.Pass + "@" + self.Host + "/" + self.Name

	s, err := mgo.Dial(dbQuery)
	if err != nil {
		panic(err)
	}

	self.S = s
	self.DB = s.DB(self.Name)
}

type DataService struct {
	DBC  *MDBC
	Post *PostService
	User *UserService
}

func (self *DataService) Init(dbc *MDBC, s *Session) {
	self.DBC = dbc
	self.Post = &PostService{
		DBC: dbc,
		S:   s,
		C:   dbc.DB.C(POST_TAB),
	}
	self.User = &UserService{
		DBC: dbc,
		S:   s,
		C:   dbc.DB.C(USER_TAB),
	}
}
