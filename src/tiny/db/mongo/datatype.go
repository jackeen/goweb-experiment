package mongo

type Post struct {
	id      string
	title   string
	content string
	auth    string
	addDate string
}

type Cate struct {
	id       string
	name     string
	explain  string
	children []string
	parent   string
}

type Tag struct {
	id      string
	name    string
	explain string
}
