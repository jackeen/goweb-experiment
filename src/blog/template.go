package main

/*
import (
//"fmt"
//"html/template"
)


type WriteContent struct {
	S string
}

func (self *WriteContent) Write(p []byte) (n int, err error) {
	self.S += string(p)
	return 0, nil
}

type TPL struct {
	TmpDir string
}

func (self *TPL) Home() {

}


type User struct {
	Name string
	Age  int
}

func main() {
	user := &User{
		Name: "tome",
		Age:  90,
	}

	c := new(WriteContent)

	t := template.New("home")
	s, _ := t.ParseFiles("../../static/default/index.html")
	s.ExecuteTemplate(c, "home", user)
	fmt.Println(c.S)
}
*/
