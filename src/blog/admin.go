package main

import (
//"encoding/json"
//"log"
)

type Admin struct {
	DBC        *MDBC
	TPL        *AdminTPL
	StaticHost string
	session    *Session
}

func (self *Admin) Init(s *Session) {
	self.session = s
}

func (self *Admin) Router(req *REQ, res *RES) {

	/*switch req.PathParm.FileName {
	case "entry":
		self.entry.Route(req, res)
	default:
		self.AdminNotFoundPage(req, res)
	}*/

}
