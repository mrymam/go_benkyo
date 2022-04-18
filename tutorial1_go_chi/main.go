package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		m := map[string]string{"message": "hello"}
		s, err := json.Marshal(m)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write(s)
	})
	http.ListenAndServe(":3000", r)
}
