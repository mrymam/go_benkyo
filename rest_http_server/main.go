package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/article", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
