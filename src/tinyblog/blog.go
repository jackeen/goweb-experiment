package main

import (
	//"flag"
	"log"
	"os/exec"
	"path"
	"strings"
)

var (
	ds         *DataService
	session    *Session
	handler    *Handler
	apiService *APIService
	admin      *Admin
	dbc        *MDBC
	staticUrl  *StaticURL
	moduleName *ModuleName
	tpl        *TplParse
	adminTpl   *TPL
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
	case moduleName.Image:
		handler.Image(req, res)
	case moduleName.API:
		apiService.Rout(req, res)
	case moduleName.Admin:
		admin.Router(req, res)
	case moduleName.Entry:
		handler.Entry(req, res)
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

	/*tpl = &TPL{
		Path:    tplDir + "blog/",
		Pattern: tplDir + "blog/*.html",
	}*/

	tpl = &TplParse{
		Path:    tplDir + "blog/",
		Pattern: tplDir + "blog/modules/*.html",
	}

	adminTpl = &TPL{
		Path:    tplDir + "admin/",
		Pattern: tplDir + "admin/*.html",
	}

	dbc = &MDBC{
		Host: "localhost",
		User: "tinyblog",
		Pass: "1234",
		Name: "tinyblog",
	}
	dbc.Init()

	session = &Session{
		Data:       make(map[string]*SessionData),
		ExpireHour: 1,
	}

	ds = &DataService{}
	ds.Init(dbc, session)

	handler = &Handler{
		Tpl:        tpl,
		StaticHost: "",
		DS:         ds,
		Session:    session,
	}

	apiService = &APIService{
		S:  session,
		DS: ds,
	}

	admin = &Admin{
		Tpl:        adminTpl,
		StaticHost: "",
		Session:    session,
		DS:         ds,
	}

	moduleName = &ModuleName{
		Home:  "",
		Post:  "post",
		Cate:  "cate",
		Date:  "date",
		Tag:   "tag",
		Image: "image",
		API:   "api",
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
