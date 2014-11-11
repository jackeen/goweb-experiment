package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("static server on :9000")
	h := http.FileServer(http.Dir("../static/"))
	log.Println(http.ListenAndServe(":9000", h))
}
