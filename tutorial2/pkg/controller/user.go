package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/onyanko-pon/go_benkyo/tutorial2/pkg/model"
	"github.com/onyanko-pon/go_benkyo/tutorial2/pkg/view"
	"golang.org/x/crypto/bcrypt"
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

type GetAllResBody struct {
	Users []view.User `json:"users"`
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
	resBody := GetAllResBody{
		Users: viewUsers,
	}
	j, err := json.Marshal(resBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

type GetOneResBody struct {
	User view.User `json:"user"`
}

func (c User) GetOne(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	fmt.Printf("hoge: %d\n", id)

	if err != nil {
		log.Print(fmt.Sprintf("invalid user id: %s", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if id <= 0 {
		log.Print(fmt.Sprintf("invalid user id (0 or less) id:%d", id))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u := model.User{}
	result := c.db.First(&u, id)

	if result.Error != nil {
		log.Print(fmt.Sprintf("db user find first: %s", result.Error.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	vu := view.User{
		ID:       u.ID,
		Username: u.Username,
	}
	resBody := GetOneResBody{
		User: vu,
	}
	j, err := json.Marshal(resBody)
	if err != nil {
		log.Print(fmt.Sprintf("json encode failed: %s", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

type CreateReqBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResBody struct {
	User view.User `json:"user"`
}

func (u User) Create(w http.ResponseWriter, r *http.Request) {

	reqBody := CreateReqBody{}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(fmt.Sprintf("request body json decode failed: %s", err.Error()))
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), 12)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(fmt.Sprintf("password encript failed: %s\n", err.Error()))
		return
	}

	user := model.User{
		ID:           0,
		Username:     reqBody.Username,
		PasswordHash: string(hash),
	}
	result := u.db.Create(&user)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(fmt.Sprintf("db user create failed :%s", result.Error.Error()))
		return
	}

	viewUser := view.User{
		ID:       user.ID,
		Username: user.Username,
	}
	resBody := CreateResBody{
		User: viewUser,
	}
	j, err := json.Marshal(resBody)

	if err != nil {
		log.Print(fmt.Sprintf("json encode failed :%s", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
