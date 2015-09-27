package main

import (
	"io"
	"log"
	"os"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type FileService struct {
	DBC *MDBC
	C   *mgo.Collection
	S   *Session
}

func (self *ImageService) Save(img *Image) *ResMessage {

}

func (self *ImageService) GetMeta() {

}

func (self *ImageService) GetFile() {

}
