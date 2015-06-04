package main

import (
	"encoding/json"
	"os"
	"reflect"
)

func LoadJson(path string, byteLen int) (map[string]interface{}, error) {

	f, _ := os.Open(path)
	defer f.Close()

	m := make(map[string]interface{})

	b := make([]byte, byteLen)
	length, readErr := f.Read(b)
	if readErr != nil {
		return m, readErr
	}
	filebyte := b[:length]
	err := json.Unmarshal(filebyte, &m)

	return m, err
}

//data format
type Format struct{}

func (self *Format) DateString(t time.Time) string {
	return t.Format(DATE_FORMAT_STR)
}

func (self *Format) O2M(o interface{}) map[string]interface{} {

	m := map[string]interface{}{}

	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)

	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		key := t.Field(i).Tag.Get("json")
		val := v.Field(i)
		if val.Type().String() == "time.Time" {
			m[key] = self.DateString(val.Interface().(time.Time))
		} else {
			m[key] = val.Interface()
		}
	}
	return m
}
