package main

import (
	//"flag"
	"log"
	"os/exec"
	"path"
	"strings"
)

const (
	DateFormatStr = "2006-01-02 15:04:05"
)

var (
	session     *Session
	handler     *Handler
	jsonService *JsonService
	admin       *Admin
	dbc         *MDBC
	staticUrl   *StaticURL
	moduleName  *ModuleName
	tpl         *TPL
	adminTpl    *TPL
)

func router(req *REQ, res *RES) {

	parm := new(UrlParmData)
	staticUrl.Parse(req.GetPath(), parm)
	req.PathParm = parm

	switch parm.Module {

	case moduleName.Home:
		handler.Index(req, res)
	case moduleName.Post:
		handler.Post(req, res)
	case moduleName.Cate:
		handler.Cate(req, res)
	case moduleName.Tag:
		handler.Tag(req, res)
	case moduleName.Date:
		handler.Date(req, res)
	case moduleName.Entry:
		handler.Entry(req, res)
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

	apppath, _ := exec.LookPath("tinyblog")
	syspath, _ := path.Split(apppath)
	baseDir := strings.TrimRight(syspath, "bin/")
	tplDir := baseDir + "/tpl/tinyblog/"
	staticDir := baseDir + "/static/"

	serverPort := "9090"
	//serverName := "localhost"
	//serverHost := serverName + ":" + serverPort

	serverIP := ""
	serverAddress := serverIP + ":" + serverPort

	tpl = new(TPL)
	tpl.Pattern = tplDir + "blog/*.html"
	adminTpl = new(TPL)
	adminTpl.Pattern = tplDir + "admin/*.html"

	dbc = &MDBC{
		Host: "localhost",
		User: "tinyblog",
		Pass: "1234",
		Name: "tinyblog",
	}
	dbc.Init()

	session = &Session{
		Data: make(map[string]*SessionData),
	}

	handler = &Handler{
		Tpl:        tpl,
		StaticHost: "",
	}
	handler.Init(dbc)

	jsonService = &JsonService{}
	jsonService.Init(dbc, session)

	admin = &Admin{
		DBC:        dbc,
		Tpl:        adminTpl,
		StaticHost: "",
	}
	admin.Init(session)

	moduleName = &ModuleName{
		Home:  "",
		Post:  "post",
		Cate:  "cate",
		Date:  "date",
		Tag:   "tag",
		Json:  "json",
		Entry: "entry",
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
