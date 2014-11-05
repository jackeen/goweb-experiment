package main

import (
	//"fmt"
	"html/template"
)

const (
	fileTail = ".html"
)

type WriteContent struct {
	Str string
}

func (self *WriteContent) Write(p []byte) (n int, err error) {
	self.Str += string(p)
	return 0, nil
}

type TPL struct {
	TmpDir string
}

func (self *TPL) Stringify(data interface{}, name string) string {

	content := new(WriteContent)
	tplFile := self.TmpDir + name + fileTail
	tpl, _ := template.New(name).ParseFiles(tplFile)
	tpl.ExecuteTemplate(content, name, data)
	return content.Str
}

func (self *TPL) PostList(data interface{}) string {
	return self.Stringify(data, "postList")
}

func (self *TPL) Post(data interface{}) string {
	return self.Stringify(data, "post")
}

//admin template page
func (self *TPL) Login(data interface{}) string {
	return self.Stringify(data, "login")
}

func (self *TPL) LoginComplete(data interface{}) string {
	return self.Stringify(data, "loginComplete")
}
