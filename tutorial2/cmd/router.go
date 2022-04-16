package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/onyanko-pon/go_benkyo/tutorial2/pkg/controller"
)

func router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	userController := controller.User{}

	r.Get("/users", userController.GetAll)

	return r
}
