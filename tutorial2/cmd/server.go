package main

import (
	"net/http"
)

func main() {
	r := router()
	http.ListenAndServe(":3000", r)
}
