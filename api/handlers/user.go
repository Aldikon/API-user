package handlers

import (
	"context"
	"net/http"
	"path"
	"strconv"

	"github.com/Aldikon/API-user/model"
)

type key int

const (
	UserIDKey key = iota
)

type user struct {
	service model.UserService
}

func NewUser(us model.UserService) *user {
	return &user{
		service: us,
	}
}

func (u *user) Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.URL.Path != "/user" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		u.get(w, r)
	case http.MethodPost:
		u.post(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (u *user) User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// TODO изменить порядок проверки, на метод важнее
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	ctx := context.WithValue(r.Context(), UserIDKey, uint(id))

	switch r.Method {
	case http.MethodGet:
		u.getByID(w, r.WithContext(ctx))
	case http.MethodDelete:
		u.delete(w, r.WithContext(ctx))
	case http.MethodPut:
		u.update(w, r.WithContext(ctx))
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
