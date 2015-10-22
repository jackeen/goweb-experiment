package main

type UserJson struct {
	S  *Session
	DS *DataService
}

func (self *UserJson) Get(req *REQ, res *RES) apiResMap {
	r := new(ResJson)
	return r.TraceMsg()
}

func (self *UserJson) Set(req *REQ, res *RES) apiResMap {
	r := new(ResJson)
	return r.TraceMsg()
}

func (self *UserJson) Put(req *REQ, res *RES) apiResMap {
	r := new(ResJson)
	return r.TraceMsg()
}

func (self *UserJson) Del(req *REQ, res *RES) apiResMap {
	r := new(ResJson)
	return r.TraceMsg()
}
