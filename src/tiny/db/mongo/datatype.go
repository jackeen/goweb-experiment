package mongo

type IdNum struct {
	Post      int
	Cate      int
	Tag       int
	User      int
	UserGroup int
}

type Post struct {
	Id           int
	Title        string
	Content      string
	Auth         int
	Cate         int
	Tags         []int
	CreateDate   string
	LastEditDate string
	EditState    bool
}

type Cate struct {
	Id       int
	Name     string
	Explain  string
	Children []int
	Parent   int
}

type Tag struct {
	Id      int
	Name    string
	Explain string
}

type UserGroup struct {
}

type User struct {
	Id          int
	Name        string
	Nick        string
	CreateDate  string
	StateNumber int
}
