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
)

func router(req *REQ, res *RES) {

	parm := new(UrlParmData)
	staticUrl.Parse(req.GetPath(), parm)
	req.PathParm = parm

	switch parm.Module {

	case moduleName.Home:
		handler.Home(req, res)
		break
	case moduleName.Post:
		handler.PostInfo(req, res)
		break
	case moduleName.Cate:
		handler.Cate(req, res)
		break
	case moduleName.Tag:
		handler.Tag(req, res)
		break
	case moduleName.Date:
		handler.Date(req, res)
		break
	case moduleName.Json:
		jsonService.GetJson(req, res)
		break
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
	staticDir := baseDir + "/static/"

	themName := "default"

	dbc = &MDBC{
		Host: "localhost",
		User: "tinyblog",
		Pass: "1234",
		Name: "tinyblog",
	}
	dbc.Init()

	handler = &Handler{
		TempLateDir: staticDir + themName + "/",
	}
	handler.Init(dbc)

	jsonService = &JsonService{}
	jsonService.Init(dbc)

	admin = &Admin{
		DBC:    dbc,
		TPLDIR: staticDir + themName + "/admin/",
	}

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
		Address: ":9090",
	}
	panic(MuxServe(httpConfig, router))

}
