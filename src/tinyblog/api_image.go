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

	//id := req.GetFormValue("id")

	return r.TraceMsg()
}
