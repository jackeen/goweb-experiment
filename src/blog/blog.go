package main

import (
	//"flag"
	"log"
	"os/exec"
	"path"
	"strings"
)

var (
	handler *Handler
	dbc     *MDBC
	URL     *StaticURL
)

func router(req *HTTPServerReq, res *HTTPServerRes) {

	parm := new(UrlParmData)
	URL.Parse(req.Path, parm)

	switch req.Path {

	case "/":
		handler.Home(req, res)
		break
	case "/post":
		handler.PostInfo(req, res)
		break
	case "/cate":
		handler.Cate(req, res)
		break
	case "/tag":
		handler.Tag(req, res)
		break
	case "/date":
		handler.Date(req, res)
	default:
		handler.NotFind(req, res)

	}
}

func main() {

	//pid := flag.Int("pid", 0, "post select id")
	//cid := flag.Int("cid", 0, "cate select id")
	//flag.Parse()

	apppath, _ := exec.LookPath("blog")
	syspath, _ := path.Split(apppath)
	baseDir := strings.TrimRight(syspath, "bin/")
	staticDir := baseDir + "/static/"

	dbc = &MDBC{
		Host: "localhost",
		User: "tinyblog",
		Pass: "1234",
		Name: "tinyblog",
	}
	dbc.Init()

	handler = &Handler{
		TempLateDir: staticDir + "default/",
	}
	handler.Init(dbc)

	URL = &StaticURL{}

	//http server start
	log.Println("The http server bind on :9090 ....\n")
	httpConfig := &HttpConfig{
		Address: ":9090",
	}
	panic(MuxServe(httpConfig, router))

}
