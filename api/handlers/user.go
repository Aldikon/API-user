package handlers

import (
	"net/http"

	"github.com/Aldikon/API-user/model"
)

type user struct {
	service model.UserService
}

func NewUser(us model.UserService) *user {
	return &user{
		service: us,
	}
}

func (u *user) User(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/user" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		u.get(w, r)
	case http.MethodPost:
		u.post(w, r)
	case http.MethodPut:
		u.update(w, r)
	case http.MethodDelete:
		u.delete(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
