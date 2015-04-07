package main

import ()

const (
	NUM_TAB    = "num"
	POST_TAB   = "post"
	CATE_TAB   = "cate"
	USER_TAB   = "user"
	TAG_TAB    = "tag"
	NAV_TAB    = "nav"
	CONFIG_TAB = "config"
)

const (
	SaveSuccess  = "save success"
	SaveFail     = "save fail"
	SaveDataFail = "save data fail"
)

type ResMessage struct {
	State   bool
	Message string
}

type ResData struct {
	State bool
	Count int
	Data  interface{}
}

//inc id num data I/O
type NumService struct{}

func (self *NumService) Init(dbc *MDBC) {
	dbc.Insert(NUM_TAB, &Num{-1, -1, -1, -1})
}

func (self *NumService) incId(dbc *MDBC, colName string, i int) *Num {
	res := &Num{}
	dbc.UpdateInc(NUM_TAB, nil, colName, i)
	dbc.SelectOne(NUM_TAB, nil, res)
	return res
}
