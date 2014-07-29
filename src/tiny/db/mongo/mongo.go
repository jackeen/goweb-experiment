package mongo

import (
	"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
)

type DBConfig struct {
	DBHost string
	DBUser string
	DBPass string
	DBName string
}

type TabName struct {
	Post   string
	Cate   string
	Tag    string
	Nav    string
	Config string
}

type DataService struct {
	DBC *DBConfig
	Tab *TabName
	s   *mgo.Session
	db  *mgo.Database
	c   *mgo.Collection
}

func (self *DataService) Close() {
	self.s.Close()
}

func (self *DataService) Connect() {

	dbc := self.DBC

	/*mongodb://user:password@hostname/dbname*/
	dbQuery := "mongodb://" + dbc.DBUser + ":" + dbc.DBPass + "@" + dbc.DBHost + "/" + dbc.DBName

	s, err := mgo.Dial(dbQuery)
	if err != nil {
		panic(err)
	}
	self.s = s
	self.db = s.DB(dbc.DBName)
}

func (self *DataService) SelTab(tabName string) {
	self.c = self.db.C(tabName)
}

func (self *DataService) Insert(data interface{}) {
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
