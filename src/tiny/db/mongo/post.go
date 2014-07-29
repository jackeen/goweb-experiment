package mongo

import (
	//"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	//"log"
	"time"
)

type PostService struct {
	DBC *DBConfig
	Tab *TabName
}

/*
func (self *PostService) Select(selector map[string]interface{}) []map[string]interface{} {

}*/

func (self *PostService) Insert(title string, content string) {

	ds := &DataService{
		DBC: self.DBC,
		Tab: self.Tab,
	}
	defer ds.Close()

	ds.Connect()
	ds.SelTab(ds.Tab.Post)

	data := &Post{
		Id_:        bson.NewObjectId(),
		Id:         1,
		Title:      title,
		Content:    content,
		CreateTime: time.Now(),
	}
	ds.Insert(data)

}

func (self *PostService) Update() {

}
