package main

import (
	"fmt"
	"net/http"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/onyanko-pon/go_benkyo/tutorial2/pkg/controller"
)

func connectDB() (gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return *db, err
}

func router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	db, err := connectDB()

	if err != nil {
		panic(err)
	}

	userController := controller.NewUserController(db)
	r.Get("/users/{id}", userController.GetOne)
	r.Get("/users", userController.GetAll)
	r.Post("/users", userController.Create)

	return r
}
