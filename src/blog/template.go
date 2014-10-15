package main

import (
	//"fmt"
	"html/template"
)

type WriteContent struct {
	Content string
}

func (self *WriteContent) Write(p []byte) (n int, err error) {
	self.Content += string(p)
	return 0, nil
}

type TPL struct {
	TmpDir string
}

func (self *TPL) PostList(data interface{}) string {
	name := "postList"
	content := new(WriteContent)
	tplFile := self.TmpDir + name + ".html"
	tpl, _ := template.New(name).ParseFiles(tplFile)
	tpl.ExecuteTemplate(content, name, data)
	return content.Content
}
