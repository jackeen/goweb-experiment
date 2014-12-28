package main

import (
	"html/template"
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

func (self *TplParse) parse(pattern string, name string, data interface{}) string {
	content := new(TPLWriteContent)
	tpl, _ := template.ParseGlob(pattern).Funcs(TPLFuncMap)
	tpl.ExecuteTemplate(content, name, data)
	return content.Str
}

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
	Pattern string
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
	Pattern string
}

func (self *AdminTPL) Home(data interface{}) string {
	return self.Stringify(data, self.Dir, "home")
}

type TplContent struct {
	TplPattern      string
	AdminTplPattern string
}

func InitTPL() (*TPL, *AdminTPL) {

}
