package main

import (
	//"flag"
	"log"
	"os/exec"
	"path"
	"strings"
)

var (
	handler     *Handler
	jsonService *JsonService
	admin       *Admin
	dbc         *MDBC
	staticUrl   *StaticURL
	moduleName  *ModuleName
	tpl         *TPL
	adminTpl    *AdminTPL
)

func router(req *REQ, res *RES) {

	parm := new(UrlParmData)
	staticUrl.Parse(req.GetPath(), parm)
	req.PathParm = parm

	switch parm.Module {

	case moduleName.Home:
		handler.Home(req, res)
	case moduleName.Post:
		handler.PostInfo(req, res)
	case moduleName.Cate:
		handler.Cate(req, res)
	case moduleName.Tag:
		handler.Tag(req, res)
	case moduleName.Date:
		handler.Date(req, res)
	case moduleName.Json:
		jsonService.GetJson(req, res)
	case moduleName.Admin:
		admin.Router(req, res)
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
	tplDir := baseDir + "/tpl/"
	staticDir := baseDir + "/static/"

	serverPort := "9090"
	serverName := "localhost"
	serverHost := serverName + ":" + serverPort
	serverIP := ""
	serverAddress := serverIP + ":" + serverPort

	tpl = &TPL{
		Dir: tplDir,
	}

	adminTpl = &AdminTPL{
		Dir: tplDir + "admin/",
	}

	dbc = &MDBC{
		Host: "localhost",
		User: "tinyblog",
		Pass: "1234",
		Name: "tinyblog",
	}
	dbc.Init()

	handler = &Handler{
		TPL: tpl,
	}
	handler.Init(dbc)

	jsonService = &JsonService{}
	jsonService.Init(dbc)

	admin = &Admin{
		DBC:           dbc,
		TPL:           adminTpl,
		StaticRootURL: serverHost + "/static/",
	}
	admin.Init()

	moduleName = &ModuleName{
		Home:  "",
		Post:  "post",
		Cate:  "cate",
		Date:  "date",
		Tag:   "tag",
		Json:  "json",
		Admin: "admin",
	}

	staticUrl = &StaticURL{}

	//http server start
	log.Println("The http server bind on :9090 ....\n")
	httpConfig := &HttpConfig{
		StaticRootDir: staticDir,
		Address:       serverAddress,
	}

	MuxServe(httpConfig, router)

}
