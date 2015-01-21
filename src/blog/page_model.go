package main

type HeadContent struct {
	PageTitle  string
	StaticHost string
}

type EntryPageData struct {
	HeadContent
	StaticHost string
}

type AdminHomeData struct {
	HeadContent
	PageTitle string
}
