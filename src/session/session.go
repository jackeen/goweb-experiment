package session

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"math/rand"
	"strconv"
	"time"
)

func MD5(s string) string {
	m := md5.New()
	io.WriteString(m, s)
	return hex.EncodeToString(m.Sum(nil))
}

func UUID() string {
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	randNum := rand.Int63()
	id := MD5(strconv.FormatInt(randNum, 16))
	return id
}

type Session struct {
	Id     string
	Expire int
	Data   map[string]interface{}
}

func (self *Session) setData(k string, v interface{}) {
	self.Datas[k] = v
}

func (self *Session) getData(k string) interface{} {
	return self.Datas[k]
}
