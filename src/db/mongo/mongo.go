package mongo

import (
	"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
)

type Config struct {
	Host string
	User string
	Pass string
	Name string
}

type TabName struct {
	Post   string
	Cate   string
	Tag    string
	Nav    string
	Config string
}

func GetTabName() *TabName {
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

func Insert(c *mgo.Collection, data interface{}) {
	err := c.Insert(data)
	if err != nil {
		panic(err)
	}
}

func Execute(dbc *Config, tabName string, fn InsertFunc, data interface{}) {

	/*mongodb://user:password@hostname/dbname*/
	dbQuery := "mongodb://" + dbc.User + ":" + dbc.Pass + "@" + dbc.Host + "/" + dbc.Name

	s, err := mgo.Dial(dbQuery)
	defer s.Close()

	if err != nil {
		panic(err)
	}

	db := s.DB(dbc.Name)
	c := db.C(tabName)

	fn(c, data)
}
