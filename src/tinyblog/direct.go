package main

func GotoAdminHome(req *REQ, res *RES) {
	res.State = 301
	res.SetHeader("Location", "/admin/home")
}

func GotoLoginErr(req *REQ, res *RES) {
	res.Response = "error"
}
