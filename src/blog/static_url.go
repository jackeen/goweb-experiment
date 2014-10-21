package main

import (
	//"log"
	"path"
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
	Module   string
	Value    string
	FileName string
	PageNum  int
}

type StaticURL struct {
}

func (self *StaticURL) Parse(reqPath string, parmData *UrlParmData) {

	filePath, fileName := path.Split(reqPath)
	parmData.FileName = fileName

	r := regexp.MustCompile(`[^/]+`)
	pathItemList := r.FindAllString(filePath, -1)

	moduleName := pathItemList[0]
	parmData.Module = moduleName

	switch moduleName {
	case POST_URL:
		self.postInfo(pathItemList, parmData)
	case CATE_URL:
		self.postListByCate(pathItemList, parmData)
	}
}

/*func (self *StaticURL) ParsePageNum(path string) string {
	r := regexp.MustCompile(`/[\d]+/`)
	s := r.FindString(path)
	return s
}*/

func (self *StaticURL) postListByCate(paths []string, parm *UrlParmData) {

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
