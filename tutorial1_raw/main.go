package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
