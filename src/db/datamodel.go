package db

type Post struct {
	Id      string
	Title   string
	Content string
	Auth    string
	AddDate string
}

type Cate struct {
	Id   string
	name string
}

type Tag struct {
	Id      string
	name    string
	explain string
}
