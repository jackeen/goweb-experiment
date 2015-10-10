package main

import (
	"io"
	//"log"
	"net/http"
	//"strings"
	"time"
)

type HttpConfig struct {
	StaticRootDir string
	Address       string
}

//request struct
type REQ struct {
	R        *http.Request
	PathParm *UrlParmData
}

func (self *REQ) Init(r *http.Request) {
	self.R = r
}

func (self *REQ) GetPath() string {
	return self.R.URL.Path
}

func (self *REQ) GetUrlParm() map[string][]string {
	return self.R.URL.Query()
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
	return self.R.FormValue(k)
}

func (self *REQ) GetCookieValues() map[string]string {
	m := make(map[string]string)
	cookies := self.R.Cookies()
	for _, v := range cookies {
		m[v.Name] = v.Value
	}
	return m
}

func (self *REQ) GetOneCookieValue(k string) string {
	cookies := self.R.Cookies()
	cookieValue := ""
	for _, v := range cookies {
		if v.Name == k {
			cookieValue = v.Value
			break
		}
	}
	return cookieValue
}

func (self *REQ) GetHeaders(k string) []string {
	return self.R.Header[k]
}

//response struct
type RES struct {
	W        http.ResponseWriter
	State    int
	Response string
}

func (self *RES) Init(w http.ResponseWriter) {
	self.W = w
}

func (self *RES) CreateCookie() *http.Cookie {
	return new(http.Cookie)
}

func (self *RES) SetCookie(c *http.Cookie) {
	http.SetCookie(self.W, c)
}

func (self *RES) DelCookie(c *http.Cookie) {
	c.Expires = time.Now().AddDate(-1, 0, 0)
	http.SetCookie(self.W, c)
}

func (self *RES) SetHeader(k string, v string) {
	self.W.Header().Set(k, v)
}

func (self *RES) SetJsonHeader() {
	self.W.Header().Set("Content-Type", "application/json;charset=UTF-8")
}

func (self *RES) SetImageHeader() {
	self.W.Header().Set("", "")
}

func (self *RES) SetDownloadHeader(name string) {
	v := "attachment; filename=" + name
	self.W.Header().Set("Content-Disposition", v)
}

//handler function type
type RouterFunc func(req *REQ, res *RES)

//
type blogHandler struct {
	Router RouterFunc
}

func (self *blogHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	dataReq := &REQ{}
	dataReq.Init(req)

	dataRes := &RES{}
	dataRes.State = 200
	dataRes.Init(w)

	//execute define blog router
	self.Router(dataReq, dataRes)

	w.WriteHeader(dataRes.State)
	io.WriteString(w, dataRes.Response)
}

func MuxServe(conf *HttpConfig, h RouterFunc) {

	bh := &blogHandler{
		Router: h,
	}
	http.Handle("/", bh)

	//static serve
	jsDir := http.Dir(conf.StaticRootDir + "js/")
	jsh := http.StripPrefix("/js/", http.FileServer(jsDir))
	http.Handle("/js/", jsh)

	styleDir := http.Dir(conf.StaticRootDir + "style/")
	styleh := http.StripPrefix("/style/", http.FileServer(styleDir))
	http.Handle("/style/", styleh)

	/*manifestDir := http.Dir(conf.StaticRootDir + "manifest/")
	manifesth := http.StripPrefix("/manifest/", http.FileServer(manifestDir))
	http.Handle("/manifest/", manifesth)*/

	http.ListenAndServe(conf.Address, nil)
}
