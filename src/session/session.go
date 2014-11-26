package session

import (
	"time"
)

type Session struct {
	Id     string
	Name   string
	Expire time.Time
	Datas  map[string]interface{}
}

func (self *Session) setData(k string, v interface{}) {
	self.Datas[k] = v
}

func (self *Session) getData(k string) interface{} {
	return self.Datas[k]
}
