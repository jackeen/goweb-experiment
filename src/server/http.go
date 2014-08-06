package main

import (
	"io"
	"net/http"

	//"net/url"
	//"path"
	//"regexp"
	//"strings"
)

type HttpConfig struct {
	Address string
}

type mux struct {
}

func (this *mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	path := req.URL.Path
	query := req.URL.Query()

	source := "Hi blog!"
	state := 200

	cookie := &http.Cookie{
		Name:  "session",
		Value: "xx-xx-xx-xx-xx",
	}
	http.SetCookie(w, cookie)

	w.WriteHeader(state)
	io.WriteString(w, source)
}

func muxServe(conf HttpConfig) error {

	m := &mux{}

	s := &http.Server{
		Addr:    conf.Address,
		Handler: m,
	}
	err := s.ListenAndServe()
	return err
}
