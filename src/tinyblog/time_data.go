package main

import (
	//"fmt"
	"labix.org/v2/mgo/bson"
	"log"
	"strconv"
	"time"
)

const (
	TIME_FORMAT_STR = "2006-01-02 15:04:05"
	DATE_FORMAT_STR = "2006-01-02"
)

//init this data must be point: &TimeData{}
type TimeData time.Time

//json interface
func (self TimeData) MarshalJSON() ([]byte, error) {

	return []byte(`"` + time.Time(self).Format(DATE_FORMAT_STR) + `"`), nil
}

//json interface
func (self TimeData) UnmarshalJSON(b []byte) error {
	timeStr, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}
	self = TimeData(time.Unix(int64(timeStr), 0))

	return nil
}

//mgo getter interface
func (self *TimeData) GetBSON() (interface{}, error) {

	log.Println("^^^^^^^^^^^^get", self)

	t := time.Time(*self)

	if t.IsZero() {
		return time.Now(), nil
	}

	return t, nil
}

//mgo setter interface
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
