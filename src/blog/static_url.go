package main

import (
	//"log"
	//Path "path"
	"regexp"
	//"strings"
)

const (
	HOME_URL = "/"
	POST_URL = "post"
	CATE_URL = "cate"
	TAG_URL  = "tag"
	DATE_URL = "date"
)

type UrlParmData struct {
	Key   string
	Value string
}

type StaticURL struct {
}

func (self *StaticURL) Parse(path string, parmData *UrlParmData) {
	r := regexp.MustCompile(`[^/]+`)
	pathItemList := r.FindAllString(path, -1)

	switch pathItemList[0] {
	case POST_URL:
		self.postInfo(pathItemList, parmData)
	}
}

/*func (self *StaticURL) ParsePageNum(path string) string {
	r := regexp.MustCompile(`/[\d]+/`)
	s := r.FindString(path)
	return s
}*/

func (self *StaticURL) postListByCate() {

}

func (self *StaticURL) postListByDate() {

}

func (self *StaticURL) postInfo(paths []string, parm *UrlParmData) {

}

/*
func main() {
	Url := new(StaticURL)

	homePath := "/2/"
	//catePath := "/cate/aa/bb/2/b"
	//datePath := "/date/2008/01/01/cate/"

	path, file := Path.Split(homePath)
	pathArr := Url.Parse(path)

	log.Println(path, file, pathArr)
}
*/
