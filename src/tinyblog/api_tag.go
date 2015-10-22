package main

type TagJson struct {
	S  *Session
	DS *DataService
}

func (self *TagJson) Get(req *REQ, res *RES) apiResMap {

	r := new(ResJson)
	ts := self.DS.Tag.GetList()
	r.State = true
	r.Data = ts
	return r.TraceData()
}

func (self *TagJson) Set(req *REQ, res *RES) apiResMap {
	r := new(ResJson)
	return r.TraceMsg()
}

func (self *TagJson) Put(req *REQ, res *RES) apiResMap {

	r := new(ResJson)

	name := req.GetFormValue("n")
	if name == "" {
		r.State = false
		r.Message = REQUIRED_DEFAULT
		return r.TraceMsg()
	}

	if self.DS.Tag.IsExist(name) {
		r.State = true
		r.Message = SAVE_SUCCESS
		return r.TraceMsg()
	}

	tag := &Tag{
		Name: name,
	}
	rs := self.DS.Tag.Save(tag)
	r.State = rs.State
	r.Message = rs.TraceMixMsg()
	return r.TraceMsg()
}

func (self *TagJson) Del(req *REQ, res *RES) apiResMap {

	r := new(ResJson)

	name := req.GetFormValue("n")
	if name == "" {
		r.State = false
		r.Message = REQUIRED_DEFAULT
		return r.TraceMsg()
	}

	if !self.DS.Tag.IsExist(name) {
		r.State = true
		r.Message = DEL_SUCCESS
		return r.TraceMsg()
	}

	rs := self.DS.Tag.Del(name)
	r.State = rs.State
	r.Message = rs.TraceMixMsg()
	return r.TraceMsg()
}
