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

var dbc *Config
var dbQuery string

func InitDBC(conf map[string]string) {
	dbc = &Config{
		Host: conf["Host"],
		User: conf["User"],
		Pass: conf["Pass"],
		Name: conf["Name"],
	}
	dbQuery = "mongodb://" + dbc.User + ":" + dbc.Pass + "@" + dbc.Host + "/" + dbc.Name
}

func connect() (*mgo.Session, *mgo.Database) {

	s, err := mgo.Dial(dbQuery)

	if err != nil {
		panic(err)
	}

	db := s.DB(dbc.Name)
	return s, db
}

func Insert(tabName string, data interface{}) {

	s, db := connect()
	defer s.Close()

	c := db.C(tabName)
	err := c.Insert(data)

	if err != nil {
		panic(err)
	}
}
