package main

import (
	"errors"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"reflect"
	"time"
)

const (
	TIME_FORMAT_STR = "2006-01-02 15:04:05"
	DATE_FORMAT_STR = "2006-01-02"
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
	DEL_SUCCESS      = "delete success"
	UPDATE_SUCCESS   = "update success"
	SAVE_FAIL        = "save fail"
	DEL_FAIL         = "delete fail"
	DEL_POWER_FAIL   = "delete power fail"
	UPDATE_FAIL      = "update fail"
	REQUIRED_DEFAULT = "required default"
	NOT_FOUND        = "not found"
)

const (
	POST_MODE_CODE = "101"
	USER_MODE_CODE = "102"
	CATE_MODE_CODE = "103"
	TAGE_MODE_CODE = "104"
	SYS_MODE_CODE  = "110"
	ALL_MODE_CODE  = "200"
)

const (
	MANAGE_USR_GROUP = "manage"
	EDITOR_USR_GROUP = "editor"
	NORMAL_USR_GROUP = "normal"
)

type ResMessage struct {
	State   bool
	Addr    string
	Message string
}

func (self *ResMessage) TraceMixMsg() string {
	return self.Addr + ":" + self.Message
}

type SelectData struct {
	Condition bson.M
	Sort      string
	Limit     int
	Start     int
	UUID      string
	GT        string
}

func createErr(msg string) error {
	return errors.New(msg)
}

func getResMessage(err error, msg string, code string) *ResMessage {

	rs := new(ResMessage)
	if err == nil {
		rs.State = true
		rs.Message = msg
	} else {
		rs.State = false
		rs.Message = err.Error()
	}
	rs.Addr = code
	return rs
}

func getUserResMessage(s bool, msg string, code string) *ResMessage {
	return &ResMessage{
		State:   s,
		Message: msg,
		Addr:    code,
	}
}

type Author struct {
	S *Session
}

func (self *Author) GetCurUsr(uuid string) (bool, *User) {

	var usr *User
	sd := self.S.Get(uuid)
	if sd != nil {
		return true, sd.U
	}
	return false, usr
}

func (self *Author) IsManager(uuid string) bool {
	b, usr := self.GetCurUsr(uuid)
	if b {
		if usr.Group == MANAGE_USR_GROUP {
			return true
		}
	} else {
		return false
	}
}

func (self *Author) IsEditor(uuid string) bool {
	b, usr := self.GetCurUsr(uuid)
	if b {
		if usr.Group == EDITOR_USR_GROUP {
			return true
		}
	} else {
		return false
	}
}

func (self *Author) IsUser(uuid string) bool {
	b, usr := self.GetCurUsr(uuid)
	if b {
		if usr.Group == NORMAL_USR_GROUP {
			return true
		}
	} else {
		return false
	}
}

func (self *Author) HasEditPost(uuid string, p *Post) bool {

	isLogin, usr := self.GetCurUsr(uuid)

	if isLogin {
		usrGroup := usr.Group
		if usrGroup == MANAGE_USR_GROUP {
			return true
		} else if usrGroup == EDITOR_USR_GROUP && p.Author == usr.Name {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (self *Author) HasSavePost(uuid string) bool {

	isLogin, usr := self.GetCurUsr(uuid)
	isM := usr.Group == MANAGE_USR_GROUP
	isE := usr.Group == EDITOR_USR_GROUP

	if isLogin && isE && isM {
		return true
	} else {
		return false
	}
}

func (self *Author) HasComment(uuid string) bool {}

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

	dbQuery := "mongodb://" +
		self.User + ":" + self.Pass + "@" +
		self.Host + "/" + self.Name

	s, err := mgo.Dial(dbQuery)
	if err != nil {
		panic(err)
	}

	self.S = s
	self.DB = s.DB(self.Name)
}

type DataService struct {
	DBC  *MDBC
	S    *Session
	Auth *Author
	Post *PostService
	User *UserService
	Cate *CateService
}

func (self *DataService) Init(dbc *MDBC, s *Session) {

	self.DBC = dbc
	self.S = s

	self.Auth = &Author{
		S: s,
	}

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
	self.Cate = &CateService{
		DBC: dbc,
		S:   s,
		C:   dbc.DB.C(CATE_TAB),
	}

}

//data format
type Format struct{}

func (self *Format) DateString(t time.Time) string {
	return t.Format(DATE_FORMAT_STR)
}

func (self *Format) O2M(o interface{}) map[string]interface{} {

	m := map[string]interface{}{}

	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)

	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		key := t.Field(i).Tag.Get("json")
		val := v.Field(i)
		if val.Type().String() == "time.Time" {
			m[key] = self.DateString(val.Interface().(time.Time))
		} else {
			m[key] = val.Interface()
		}
	}
	return m
}

//split page module
/*
type SplitPageCache struct {
	pageIndexMap map[int]time.Time
	timer        *time.Ticker
}

func (self *SplitPageCache) reset(c <-chan time.Time) {
	<-c
	self.pageIndexMap = make(map[int]time.Time)
}

func (self *SplitPageCache) Init() {
	self.timer = time.NewTicker(3 * time.Minute)
	go self.reset(self.timer.C)
}

func (self *SplitPageCache) Stop() {
	self.timer.Stop()
}

func (self *SplitPageCache) Add(i int, lastTime time.Time) {
	self.pageIndexMap[i] = lastTime
}

func (self *SplitPageCache) Get(i int) (time.Time, bool) {
	t := self.pageIndexMap[i]
	s := true
	//time zero
	return t, s
}
*/
