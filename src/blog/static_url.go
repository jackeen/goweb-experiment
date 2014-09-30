package main

import (
	"log"
	Path "path"
	"regexp"
	//"strings"
)

const (
	HOME_URL = "/"
	POST_URL = "/post"
	CATE_URL = "/cate"
	TAG_URL  = "/tag"
	DATE_URL = "/date"
)

type StaticURL struct {
}

func (self *StaticURL) Parse(path string) []string {

	reg := regexp.MustCompile(`[^/]+`)
	s := reg.FindAllString(path, -1)

	return s
}

func (self *StaticURL) ParsePageNum(path string) string {
	r := regexp.MustCompile(`/[\d]+/`)
	s := r.FindString(path)
	return s
}

func (self *StaticURL) Route(path string) {

}

func (self *StaticURL) getCate() {

}

func (self *StaticURL) getDate() {

}

func (self *StaticURL) getPost() {

}

func main() {
	url := new(StaticURL)

	catePath := "/cate/aa/bb/2/"
	//datePath := "/date/2008/01/01/cate/"

	path, file := Path.Split(catePath)

	pathArr := url.Parse(path)

	url.Route(path)

	log.Println(path, file, pathArr)

}
