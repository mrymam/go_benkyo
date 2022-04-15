package main

import (
	"bytes"
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
	case "POST":
		postArticle(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	article := Article{
		Title: "title",
	}
	articles := []Article{article}
	s, _ := json.Marshal(articles)
	fmt.Fprintf(w, string(s))
}

func postArticle(w http.ResponseWriter, r *http.Request) {
	article := Article{}

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	jsonString := buf.String()
	err := json.Unmarshal([]byte(jsonString), &article)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s, _ := json.Marshal(article)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(s))
}
