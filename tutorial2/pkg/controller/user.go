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

func (u User) Create(w http.ResponseWriter, r *http.Request) {

	user := model.User{
		ID:           0,
		Username:     "hogehoge",
		PasswordHash: "hugahuga",
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
