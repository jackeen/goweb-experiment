package main

import (
	"io"
	"net/http"
)

type HttpConfig struct {
	Address string
}

type HTTPServerReq struct {
	Path     string
	PathParm *UrlParmData
	Query    map[string][]string
	Headers  map[string][]string
	Cookies  map[string]string
}

type HTTPServerRes struct {
	State    int
	Cookies  map[string]string
	Headers  map[string]string
	Response string
}

type HttpServerHandler func(req *HTTPServerReq, res *HTTPServerRes)

type mux struct {
	ServerHandler HttpServerHandler
}

func (self *mux) setCookies(w http.ResponseWriter, c map[string]string) {
	var cookie *http.Cookie
	for k, v := range c {
		cookie = &http.Cookie{
			Name:  k,
			Value: v,
		}
		http.SetCookie(w, cookie)
	}
}

func (self *mux) parseCookie(c []*http.Cookie) map[string]string {
	m := make(map[string]string)
	for _, v := range c {
		m[v.Name] = v.Value
	}
	return m
}

func (self *mux) setHeaders(w http.ResponseWriter, h map[string]string) {
	for i, v := range h {
		w.Header().Set(i, v)
	}
}

func (self *mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	path := req.URL.Path
	query := req.URL.Query()
	cookies := req.Cookies()

	dataReq := &HTTPServerReq{
		Path:    path,
		Query:   query,
		Cookies: self.parseCookie(cookies),
		Headers: req.Header,
	}
	dataRes := &HTTPServerRes{}

	self.ServerHandler(dataReq, dataRes)
	self.setCookies(w, dataRes.Cookies)
	self.setHeaders(w, dataRes.Headers)

	w.WriteHeader(dataRes.State)
	io.WriteString(w, dataRes.Response)
}

func MuxServe(conf *HttpConfig, h HttpServerHandler) error {

	m := &mux{
		ServerHandler: h,
	}

	s := &http.Server{
		Addr:    conf.Address,
		Handler: m,
	}
	err := s.ListenAndServe()
	return err
}
