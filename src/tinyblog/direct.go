package main

const (
	URL_LOGIN      = "/entry/login"
	URL_ADMIN_HOME = "/admin/home"
)

func GotoAdminHome(req *REQ, res *RES) {
	res.State = 301
	res.SetHeader("Location", URL_ADMIN_HOME)
}

func GotoLoginErr(req *REQ, res *RES) {
	res.Response = "error"
}
