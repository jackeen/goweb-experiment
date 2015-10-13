package main

import (
	//"io"
	"log"
	//"os"
	"strings"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type ImageService struct {
	DBC *MDBC
	C   *mgo.Collection
	FS  *mgo.GridFS
	S   *Session
}

func (self *ImageService) Save(fileName string, data []byte) *ResMessage {

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
		Size:        size,
		Cate:        "",
		CreateTime:  ct,
		EditTime:    ct,
	}

	err = self.C.Insert(img)
	return getResMessage(err, SAVE_SUCCESS, IMG_MODE_CODE)
}

func (self *ImageService) GetOne(id string) *Image {

	img := &Image{}
	self.C.FindId(bson.ObjectIdHex(id)).One(img)
	return img
}

/*func (self *ImageService) GetList() []Image {

}*/

func (self *ImageService) GetFile(name string) ([]byte, int, *ImageMeta) {

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
