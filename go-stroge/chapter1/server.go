package main

import (
	"chapter1/objects"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:12345", nil))
}
