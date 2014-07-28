package mongo

import (
	//"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
	"log"
)

func PostInsert(ds *DataService, data *Post) {
	ds.Connect()
	defer ds.Close()

	ds.SelTab(ds.PostTab)
	ds.Insert(data)

}
