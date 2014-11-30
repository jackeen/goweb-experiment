package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"math/rand"
	"strconv"
	"time"
)

const (
	SessionExpire = 10
)

type SessionData struct {
	User string
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
	Data map[string]*SessionData
}

func (self *Session) del(id string) {
	time.Sleep(SessionExpire * time.Second)
	delete(self.Data, id)
}

func (self *Session) New(sd *SessionData) string {
	id := CreateUUID()
	self.Data[id] = sd
	go self.del(id)
	return id
}

func (self *Session) Set(k string, v *SessionData) {
	self.Data[k] = v
}

func (self *Session) Get(k string) *SessionData {
	return self.Data[k]
}
