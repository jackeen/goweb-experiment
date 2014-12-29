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
type TPL struct {
	Pattern string
}

func (self *TPL) Parse(name string, data interface{}) string {
	w := new(TPLWriteContent)
	tpl, _ := template.ParseGlob(self.Pattern)
	tpl = tpl.Funcs(TPLFuncMap)
	tpl.ExecuteTemplate(w, name, data)
	return w.Str
}
