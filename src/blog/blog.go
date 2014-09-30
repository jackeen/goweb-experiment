package main

import (
	//"flag"
	"log"
)

var (
	handler *Handler
	dbc     *MDBC
)

func router(req *HTTPServerReq, res *HTTPServerRes) {

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

	dbc = &MDBC{
		Host: "localhost",
		User: "tinyblog",
		Pass: "1234",
		Name: "tinyblog",
	}
	dbc.Init()

	handler = &Handler{}
	handler.Init(dbc)

	//http server start
	log.Println("The http server start ....\n")
	httpConfig := &HttpConfig{
		Address: ":9090",
	}
	panic(MuxServe(httpConfig, router))

}
