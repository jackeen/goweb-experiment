package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"math/rand"
	"strconv"
	"time"
)

type SessionData struct {
	U     *User
	Timer *time.Timer
}

func CreateUUID() string {

	nano := time.Now().UnixNano()
	rand.Seed(nano)
	randNum := rand.Int63()

	m := md5.New()
	io.WriteString(m, strconv.FormatInt(randNum, 16))

	return hex.EncodeToString(m.Sum(nil))
}

type Session struct {
	ExpireHour time.Duration
	Data       map[string]*SessionData
}

func (self *Session) del(id string, c <-chan time.Time) {
	<-c
	delete(self.Data, id)
}

func (self *Session) getExpire() time.Duration {
	return self.ExpireHour * time.Hour
}

func (self *Session) New(sd *SessionData) string {
	id := CreateUUID()
	sd.Timer = time.NewTimer(self.getExpire())
	self.Data[id] = sd
	go self.del(id, sd.Timer.C)
	return id
}

func (self *Session) Destroy(uuid string) bool {
	sd := self.Data[uuid]
	if sd == nil {
		return false
	} else {
		sd.Timer.Reset(0)
		return true
	}
}

func (self *Session) reFresh(sd *SessionData) {
	sd.Timer.Reset(self.getExpire())
}

func (self *Session) Set(uuid string, sd *SessionData) {
	self.Data[uuid] = sd
}

func (self *Session) Get(uuid string) *SessionData {
	return self.Data[uuid]
}

func (self *Session) GetCurUsr(uuid string) (bool, *User) {

	var usr *User
	sd := self.Get(uuid)
	if sd != nil {
		return true, sd.U
	}
	return false, usr
}

func (self *Session) IsLogin(uuid string) bool {
	u := self.Get(uuid)
	if u != nil {
		self.reFresh(u)
		return true
	} else {
		return false
	}
}
