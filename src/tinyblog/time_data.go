package main

import (
	//"fmt"
	"labix.org/v2/mgo/bson"
	//"log"
	"time"
)

const (
	TIME_FORMAT_STR = "2006-01-02 15:04:05"
	DATE_FORMAT_STR = "2006-01-02"
)

type TimeData time.Time

func (self TimeData) MarshalJSON() ([]byte, error) {

	return []byte(`"` + time.Time(self).Format(DATE_FORMAT_STR) + `"`), nil
}

/*func (self TimeData) UnmarshalJSON(b []byte) error {

}*/

func (self *TimeData) GetBSON() (interface{}, error) {

	//log.Println("^^^^^^^^^^^^get", self)

	if time.Time(*self).IsZero() {
		return nil, nil
	}

	return time.Time(*self), nil
}

func (self *TimeData) SetBSON(raw bson.Raw) error {

	var t time.Time

	if err := raw.Unmarshal(&t); err != nil {
		return err
	}

	*self = TimeData(t)

	//log.Println("^^^^^^^^^^^^set", self)

	return nil
}

func (self *TimeData) DateString() string {
	return time.Time(*self).Format(DATE_FORMAT_STR)
}

func (self *TimeData) TimeString() string {
	return time.Time(*self).Format(TIME_FORMAT_STR)
}
