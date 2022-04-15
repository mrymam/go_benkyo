package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	Title string `json:"title"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getArticle(w, r)
		break
	case "POST":
		postArticle(w, r)
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	article := Article{
		Title: "title",
	}
	s, _ := json.Marshal(article)
	fmt.Fprintf(w, string(s))
}

func postArticle(w http.ResponseWriter, r *http.Request) {

}
