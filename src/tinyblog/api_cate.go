package main

//cate api
type CateJson struct {
	S  *Session
	DS *DataService
}

func (self *CateJson) Get(req *REQ, res *RES) apiResMap {

	r := new(ResJson)
	qParent := req.GetFormValue("p")

	cs := self.DS.Cate.GetNames(qParent)
	r.State = true
	r.Data = cs
	r.Count = len(cs)
	return r.TraceListData()
}

func (self *CateJson) Set(req *REQ, res *RES) apiResMap {
	r := new(ResJson)
	return r.TraceMsg()
}

func (self *CateJson) Put(req *REQ, res *RES) apiResMap {

	r := new(ResJson)
	qName := req.GetFormValue("n")
	qParent := req.GetFormValue("p")

	if qName == "" {
		r.State = false
		r.Message = REQUIRED_DEFAULT
		return r.TraceMsg()
	}

	c := &Cate{
		Name:   qName,
		Parent: qParent,
	}
	rs := self.DS.Cate.Save(c)

	r.State = rs.State
	r.Message = rs.TraceMixMsg()
	return r.TraceMsg()
}

func (self *CateJson) Del(req *REQ, res *RES) apiResMap {
	r := new(ResJson)
	return r.TraceMsg()
}
