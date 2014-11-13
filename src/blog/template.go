package main

import (
	//"fmt"
	"html/template"
)

const (
	FILE_TAIL = ".html"
)

type writeContent struct {
	Str string
}

func (self *writeContent) Write(p []byte) (n int, err error) {
	self.Str += string(p)
	return 0, nil
}

//
type TplParse struct{}

func (self *TplParse) Stringify(data interface{}, p string, n string) string {
	content := new(writeContent)
	tplFile := p + n + FILE_TAIL
	tpl, _ := template.New(n).ParseFiles(tplFile)
	tpl.ExecuteTemplate(content, n, data)
	return content.Str
}

//blog page template
type TPL struct {
	TplParse
	Dir string
}

func (self *TPL) PostList(data interface{}) string {
	return self.Stringify(data, self.Dir, "postList")
}

func (self *TPL) Post(data interface{}) string {
	return self.Stringify(data, self.Dir, "post")
}

//admin template page
type AdminTPL struct {
	TplParse
	Dir string
}

func (self *AdminTPL) Login(data interface{}) string {
	return self.Stringify(data, self.Dir, "login")
}

func (self *AdminTPL) LoginComplete(data interface{}) string {
	return self.Stringify(data, self.Dir, "loginComplete")
}
