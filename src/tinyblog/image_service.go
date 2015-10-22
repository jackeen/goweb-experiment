package main

import (
	//"io"
	"log"
	//"os"
	"strconv"
	"strings"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type ImageService struct {
	DBC   *MDBC
	C     *mgo.Collection
	CateC *mgo.Collection
	FS    *mgo.GridFS
	S     *Session
}

func (self *ImageService) SaveImg(fileName string, data []byte) *ResMessage {

	id := bson.NewObjectId()
	ct := time.Now()

	nameSub := strings.Split(fileName, ".")
	name := nameSub[0]
	extName := nameSub[1]

	gf, err := self.FS.Create(id.Hex())
	if err != nil {
		log.Println(err)
	}

	size, err := gf.Write(data)
	if err != nil {
		log.Println(err)
	}
	gf.SetMeta(&ImageMeta{
		ContentName: extName,
		Name:        name,
	})
	defer gf.Close()

	if err != nil {
		return getUserResMessage(false, err.Error(), IMG_MODE_CODE)
	}

	img := &Image{
		Id_:         id,
		FileName:    id.Hex(),
		Name:        name,
		ContentName: extName,
		Size:        strconv.Itoa(size),
		Cate:        "",
		CreateTime:  ct,
		EditTime:    ct,
	}

	err = self.C.Insert(img)
	return getResMessage(err, SAVE_SUCCESS, IMG_MODE_CODE)
}

func (self *ImageService) DelImg(id string) *ResMessage {

	if !bson.IsObjectIdHex(id) {
		return getUserResMessage(false, NOT_ID, IMG_MODE_CODE)
	}

	err := self.FS.Remove(id)
	if err != nil {
		return getUserResMessage(false, DEL_FAIL, IMG_MODE_CODE)
	}

	err = self.C.RemoveId(bson.ObjectIdHex(id))
	return getResMessage(err, DEL_SUCCESS, IMG_MODE_CODE)
}

func (self *ImageService) GetImgList(cateName string) []Image {

	q := self.C.Find(bson.M{"cate": cateName})

	n, err := q.Count()
	if err != nil {
		n = 0
	}

	imgList := make([]Image, n)
	if n > 0 {
		q.All(&imgList)
	}

	return imgList
}

func (self *ImageService) GetImgFile(name string) ([]byte, int, *ImageMeta) {

	gf, err := self.FS.Open(name)
	if err != nil {
		log.Println(err)
	}
	defer gf.Close()

	b := make([]byte, gf.Size())
	size, err := gf.Read(b)
	if err != nil {
		log.Println(err)
	}

	imgMeta := &ImageMeta{}
	err = gf.GetMeta(imgMeta)
	if err != nil {
		log.Println(err)
	}

	return b, size, imgMeta
}

func (self *ImageService) SaveCate(cate *ImageCate) *ResMessage {

	t := time.Now()

	cate.Id_ = bson.NewObjectId()
	cate.CreateTime = t
	cate.EditTime = t

	err := self.CateC.Insert(cate)
	return getResMessage(err, SAVE_SUCCESS, IMG_CATE_MODE_CODE)
}

func (self *ImageService) EditCate(id string, cate *ImageCate) *ResMessage {

	if !bson.IsObjectIdHex(id) {
		return getUserResMessage(false, NOT_ID, IMG_CATE_MODE_CODE)
	}

	hexId := bson.ObjectIdHex(id)
	curCate := &ImageCate{}

	err := self.CateC.FindId(hexId).One(curCate)
	if err != nil {
		return getUserResMessage(false, TARGET_NOT_EXIST, IMG_CATE_MODE_CODE)
	}

	if cate.Name != "" {
		curCate.Name = cate.Name
	}

	if cate.Explain != "" {
		curCate.Explain = cate.Explain
	}

	err = self.CateC.UpdateId(id, curCate)
	return getResMessage(err, UPDATE_SUCCESS, IMG_CATE_MODE_CODE)

}

func (self *ImageService) DelCate(id string) *ResMessage {

	if !bson.IsObjectIdHex(id) {
		return getUserResMessage(false, NOT_ID, IMG_CATE_MODE_CODE)
	}

	err := self.CateC.RemoveId(bson.ObjectIdHex(id))
	return getResMessage(err, DEL_FAIL, IMG_CATE_MODE_CODE)

}

func (self *ImageService) GetCateList() []ImageCate {

	req := self.CateC.Find(nil)

	n, err := req.Count()
	if err != nil {
		n = 0
	}

	cateList := make([]ImageCate, n)

	if n > 0 {
		req.All(cateList)

	}
	return cateList
}
