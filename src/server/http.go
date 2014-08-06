package main

import (
	"io"
	"net/http"

//"net/url"
//"path"
//"regexp"
//"strings"
)

type mux struct {
	r *router.Router
}

func (this *mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	path := req.URL.Path
	query := req.URL.Query()

	source, state := this.r.Rout(path, query)

	cookie := &http.Cookie{
		Name:  "session",
		Value: "xx",
	}
	http.SetCookie(w, cookie)

	w.WriteHeader(state)
	io.WriteString(w, source)
}

func muxServe(httpConf map[string]string) error {

	m := &mux{}

	m.r = &router.Router{}

	s := &http.Server{
		Addr:    httpConf["addr"],
		Handler: m,
	}
	err := s.ListenAndServe()
	return err
}
