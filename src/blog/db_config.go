package main

var DBConfig map[string]string = map[string]string{
	"Host": "localhost",
	"User": "tinyblog",
	"Pass": "1234",
	"Name": "tinyblog",
}

var DBTable map[string]string = map[string]string{
	"Post":   "post",
	"Cate":   "cate",
	"Tag":    "tag",
	"Nav":    "nav",
	"Config": "config",
}
