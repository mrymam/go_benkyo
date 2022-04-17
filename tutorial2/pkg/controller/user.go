package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/onyanko-pon/go_benkyo/tutorial2/pkg/model"
	"github.com/onyanko-pon/go_benkyo/tutorial2/pkg/view"
	"gorm.io/gorm"
)

type User struct {
	db gorm.DB
}

func NewUserController(db gorm.DB) User {
	return User{
		db: db,
	}
}

func (u User) GetAll(w http.ResponseWriter, r *http.Request) {

	users := []model.User{}
	result := u.db.Find(&users)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	viewUsers := []view.User{}
	for _, u := range users {
		vu := view.User{
			ID:       u.ID,
			Username: u.Username,
		}
		viewUsers = append(viewUsers, vu)
	}

	j, err := json.Marshal(viewUsers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

type CreateBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func

func (u User) Create(w http.ResponseWriter, r *http.Request) {

	body := CreateBody{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(fmt.Printf("request body json decode failed :%s", err.Error()))
		return
	}

	hash :=[]byte("")
	err = bcrypt.CompareHashAndPassword(hash,[]byte(body.Password))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(fmt.Printf("password encript failed :%s", err.Error()))
		return
	}

	user := model.User{
		ID:           0,
		Username:     body.Username,
		PasswordHash: hash,
	}
	result := u.db.Create(&user)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(fmt.Printf("db user create failed :%s", result.Error.Error()))
		return
	}

	viewUser := view.User{
		ID:       user.ID,
		Username: user.Username,
	}
	j, err := json.Marshal(viewUser)

	if err != nil {
		log.Print(fmt.Printf("json encode failed :%s", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
