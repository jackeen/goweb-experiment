package main

import (
	"fmt"
	"tiny/db/mongo"
)

func main() {
	mongo.Init()
	fmt.Println("haha")
}
