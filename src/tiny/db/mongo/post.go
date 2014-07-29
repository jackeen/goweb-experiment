package mongo

import (
	//"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	//"log"
	"time"
)

func PostInsert(dbc *DBConfig, title string, content string) {

	data := &Post{
		Id_:        bson.NewObjectId(),
		Id:         1,
		Title:      title,
		Content:    content,
		CreateTime: time.Now(),
	}

	DBConnect(dbc, DBInsert, getTabName().Post, data)

}
