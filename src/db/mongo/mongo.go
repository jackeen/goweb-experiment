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

func getTabName() *TabName {
	tableName := &TabName{
		Post:   "post",
		Cate:   "cate",
		Tag:    "tag",
		Nav:    "nav",
		Config: "config",
	}
	return tableName
}

type InsertFunc func(*mgo.Collection, interface{})

func DBInsert(c *mgo.Collection, data interface{}) {
	err := c.Insert(data)
	if err != nil {
		panic(err)
	}
}

func DBConnect(dbc *DBConfig, fn InsertFunc, tabName string, data interface{}) {

	/*mongodb://user:password@hostname/dbname*/
	dbQuery := "mongodb://" + dbc.DBUser + ":" + dbc.DBPass + "@" + dbc.DBHost + "/" + dbc.DBName

	s, err := mgo.Dial(dbQuery)
	defer s.Close()

	if err != nil {
		panic(err)
	}

	db := s.DB(dbc.DBName)
	c := db.C(tabName)

	fn(c, data)
}
