package main

import (
	"bytes"
	"html/template"
	"log"
)

var TPLFuncMap = template.FuncMap{
	"attr": func(s string) template.HTMLAttr {
		return template.HTMLAttr(s)
	},
	"html": func(s string) template.HTML {
		return template.HTML(s)
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
	Path    string
	Pattern string
}

func (self *TPL) Parse(name string, data interface{}) string {
	w := new(TPLWriteContent)
	tpl := template.New(name).Funcs(TPLFuncMap)
	tpl.ParseGlob(self.Pattern)
	tpl.Execute(w, data)
	return w.Str
}

//
type TplParse struct {
	Path    string
	Pattern string
}

func (self *TplParse) Parse(res *RES, f string, d PageData) {

	buf := &bytes.Buffer{}

	tpl, err := template.ParseFiles(self.Path + "warp.html")
	if err != nil {
		log.Println(err)
		return
	}

	tpl, err = tpl.ParseGlob(self.Pattern)
	if err != nil {
		log.Println(err)
		return
	}

	tpl, err = tpl.Funcs(TPLFuncMap).ParseFiles(self.Path + f)
	if err != nil {
		log.Println(err)
		return
	}

	err = tpl.Execute(buf, d)
	if err != nil {
		log.Println(err)
		return
	}

	res.Response = buf.String()
}
