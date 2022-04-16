package controller

import (
	"encoding/json"
	"net/http"

	"github.com/onyanko-pon/go_benkyo/tutorial2/pkg/model"
	"github.com/onyanko-pon/go_benkyo/tutorial2/pkg/view"
)

type User struct{}

func (u User) GetAll(w http.ResponseWriter, r *http.Request) {
	user := model.User{
		ID:       1,
		Username: "hogehoge",
	}

	viewUser := view.User{
		ID:       user.ID,
		Username: user.Username,
	}

	users := []view.User{viewUser}

	j, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
