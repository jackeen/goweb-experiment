package mongo

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type DataService struct {
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	PostTab   string
	CateTab   string
	TagTab    string
	NavTab    string
	ConfigTab string
	s         *mgo.Session
	db        *mgo.Database
	c         *mgo.Collection
}

func (self *DataService) Close() {
	self.S.Close()
}

func (self *DataService) Connect() {

	/*mongodb://user:password@hostname/dbname*/
	dbQuery := "mongodb://" + self.DBUser + ":" + self.DBPass + "@" + self.DBHost + "/" + self.DBName

	s, err := mgo.Dial(dbQuery)

	if err != nil {
		panic(err)
	}
	self.s = s
	self.db = s.DB(self.DBName)
}

func (self *DataService) SelTab(tabName string) {
	self.c = self.db.C(tabName)
}

func (self *DataService) Insert(data *map[string]interface{}) {
	c := self.c
	err := c.Insert(data)
	if err != nil {
		panic(err)
	}
}

func (self *DataService) Update() {

}

func (self *DataService) Select() {

}
