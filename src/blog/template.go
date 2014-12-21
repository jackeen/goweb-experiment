package main

import (
	//"fmt"
	"html/template"
)

const (
	FILE_TAIL = ".html"
)

var TPLFuncMap = template.FuncMap{
	"attr": func(s string) template.HTMLAttr {
		return template.HTMLAttr(s)
	},
}

type TPLWriteContent struct {
	Str string
}

func (self *TPLWriteContent) Write(p []byte) (n int, err error) {
	self.Str += string(p)
	return 0, nil
}

//
type TplParse struct{}

func (self *TplParse) Stringify(data interface{}, p string, n string) string {
	content := new(TPLWriteContent)
	tplFile := p + n + FILE_TAIL
	tpl, _ := template.New(n).Funcs(TPLFuncMap).ParseFiles(tplFile)
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

func (self *TPL) Login(data interface{}) string {
	return self.Stringify(data, self.Dir, "login")
}

//admin template page
type AdminTPL struct {
	TplParse
	Dir string
}

func (self *AdminTPL) LoginComplete(data interface{}) string {
	return self.Stringify(data, self.Dir, "loginComplete")
}
