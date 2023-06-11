package main

import (
	"chapter1/locate"
	"chapter1/objects"
	"log"
	"net/http"
)

func main() {
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:12366", nil))
}
