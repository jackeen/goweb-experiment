package main

import (
	"io"
	//"log"
	"net/http"
)

type HttpConfig struct {
	Address string
}

//request struct
type REQ struct {
	r        *http.Request
	PathParm *UrlParmData
}

func (self *REQ) Init(r *http.Request) {
	self.r = r
}

func (self *REQ) GetPath() string {
	return self.r.URL.Path
}

func (self *REQ) GetUrlParm() map[string][]string {
	return self.r.URL.Query()
}

func (self *REQ) GetUrlOneValue(k string) string {

	m := self.GetUrlParm()
	vals := m[k]

	if len(vals) > 0 {
		return vals[0]
	} else {
		return ""
	}
}

func (self *REQ) GetFormValue(k string) string {
	return self.r.FormValue(k)
}

func (self *REQ) GetCookies() map[string]string {
	m := make(map[string]string)
	cookies := self.r.Cookies()
	for _, v := range cookies {
		m[v.Name] = v.Value
	}
	return m
}

func (self *REQ) GetHeaders(k string) []string {
	return self.r.Header[k]
}

//response struct
type RES struct {
	w        http.ResponseWriter
	State    int
	Response string
}

func (self *RES) Init(w http.ResponseWriter) {
	self.w = w
}

func (self *RES) SetCookie(k string, v string) {
	c := &http.Cookie{
		Name:  k,
		Value: v,
	}
	http.SetCookie(self.w, c)
}

func (self *RES) SetHeader(k string, v string) {
	self.w.Header().Set(k, v)
}

//handler function type
type HandlerFunc func(req *REQ, res *RES)

//
type blogHandler struct {
	Handler HandlerFunc
}

func (self *blogHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	dataReq := &REQ{}
	dataReq.Init(req)

	dataRes := &RES{}
	dataRes.Init(w)

	//execute define handler function
	self.Handler(dataReq, dataRes)

	w.WriteHeader(dataRes.State)
	io.WriteString(w, dataRes.Response)
}

func MuxServe(conf *HttpConfig, h HandlerFunc) {

	bh := &blogHandler{
		Handler: h,
	}
	http.Handle("/", bh)

	fh := http.FileServer(http.Dir("../static"))
	http.Handle("/static/", fh)

	http.ListenAndServe(conf.Address, nil)
}
