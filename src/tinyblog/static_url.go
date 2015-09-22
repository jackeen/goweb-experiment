package main

import (
	//"log"
	"path"
	"regexp"
	//"strings"
)

/*
const (
	HOME = ""
	POST = "post"
	CATE = "cate"
	TAG  = "tag"
	DATE = "date"
)
*/

type UrlParmData struct {
	PathItems []string
	Module    string
	Value     string
	FileName  string
	PageNum   int
}

type ModuleName struct {
	Home  string
	Post  string
	Cate  string
	Date  string
	Tag   string
	File  string
	Entry string
	API   string
	Admin string
}

type StaticURL struct {
}

func (self *StaticURL) Parse(reqPath string, parmData *UrlParmData) {

	filePath, fileName := path.Split(reqPath)
	parmData.FileName = fileName

	r := regexp.MustCompile(`[^/]+`)
	pathItemList := r.FindAllString(filePath, -1)

	var moduleName string
	if len(pathItemList) > 0 {
		moduleName = pathItemList[0]
	} else {
		moduleName = ""
	}

	parmData.Module = moduleName
	parmData.PathItems = pathItemList
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
