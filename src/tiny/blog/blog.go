package main

import (
	"fmt"
	mdb "tiny/db/mongo"
)

func main() {
	mdb.Init()
	fmt.Println("haha")
}
