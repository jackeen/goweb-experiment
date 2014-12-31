package main

type HeadContent struct {
	PageTitle string
}

type EntryPageData struct {
	HeadContent
	StaticHost string
}

type AdminHomeData struct {
	HeadContent
	PageTitle string
}
