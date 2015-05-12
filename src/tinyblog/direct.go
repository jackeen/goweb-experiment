package main

const (
	URL_LOGIN      = "/entry/"
	URL_ADMIN_HOME = "/admin/home"
	URL_HOME       = "/"
)

func GotoHome(req *REQ, res *RES) {
	res.State = 301
	res.SetHeader("Location", URL_HOME)
}

func GotoAdminHome(req *REQ, res *RES) {
	res.State = 301
	res.SetHeader("Location", URL_ADMIN_HOME)
}

func GotoLogin(req *REQ, res *RES) {
	res.State = 301
	res.SetHeader("Location", URL_LOGIN)
}

func GotoLoginErr(req *REQ, res *RES) {
	res.Response = "error"
}
