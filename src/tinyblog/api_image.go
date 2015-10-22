package main

import (
	"io/ioutil"
)

type ImageAPI struct {
	S  *Session
	DS *DataService
}

func (self *ImageAPI) Get(req *REQ, res *RES) apiResMap {

	r := new(ResJson)

	imgList := self.DS.Img.GetImgList("")
	r.State = true
	r.Count = len(imgList)
	r.Data = self.DS.F.TransImageList(imgList)

	return r.TraceListData()
}

func (self *ImageAPI) Set(req *REQ, res *RES) apiResMap {

	r := new(ResJson)

	return r.TraceMsg()
}

func (self *ImageAPI) Put(req *REQ, res *RES) apiResMap {

	r := new(ResJson)

	file, header, err := req.R.FormFile("photo")
	defer file.Close()

	if err != nil {
		return r.TraceMsg()
	}

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		return r.TraceMsg()
	}

	rs := self.DS.Img.SaveImg(header.Filename, bytes)

	r.State = rs.State
	r.Message = rs.TraceMixMsg()

	return r.TraceMsg()
}
func (self *ImageAPI) Del(req *REQ, res *RES) apiResMap {

	r := new(ResJson)

	id := req.GetFormValue("id")

	if id == "" {
		r.State = false
		r.Message = REQUIRED_DEFAULT
	} else {
		rs := self.DS.Img.DelImg(id)
		r.State = rs.State
		r.Message = rs.TraceMixMsg()
	}

	return r.TraceMsg()
}

type ImgCateApi struct {
	S  *Session
	DS *DataService
}

func (self *ImgCateApi) Get(req *REQ, res *RES) apiResMap {

	r := new(ResJson)

	cateList := self.DS.Img.GetCateList()
	r.State = true
	r.Count = len(cateList)
	r.Data = self.DS.F.TransImgCateList(cateList)

	return r.TraceMsg()
}

func (self *ImgCateApi) Set(req *REQ, res *RES) apiResMap {
	r := new(ResJson)
	return r.TraceMsg()
}

func (self *ImgCateApi) Put(req *REQ, res *RES) apiResMap {

	r := new(ResJson)

	name := req.GetFormValue("name")
	exp := req.GetFormValue("exp")
	if name == "" || exp == "" {
		r.State = false
		r.Message = REQUIRED_DEFAULT
	} else {
		c := &ImageCate{
			Name:    name,
			Explain: exp,
		}
		rs := self.DS.Img.SaveCate(c)
		r.State = rs.State
		r.Message = rs.TraceMixMsg()
	}

	return r.TraceMsg()
}

func (self *ImgCateApi) Del(req *REQ, res *RES) apiResMap {

	r := new(ResJson)

	id := req.GetFormValue("id")
	if id == "" {
		r.State = false
		r.Message = REQUIRED_DEFAULT
	} else {
		rs := self.DS.Img.DelCate(id)
		r.State = rs.State
		r.Message = rs.TraceMixMsg()
	}

	return r.TraceMsg()
}
