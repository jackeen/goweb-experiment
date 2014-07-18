package mongo

type Post struct {
	Id      string
	Title   string
	Content string
	Auth    string
	AddDate string
}

type Cate struct {
	Id       string
	Name     string
	Explain  string
	Children []string
	Parent   string
}

type Tag struct {
	Id      string
	Name    string
	Explain string
}
