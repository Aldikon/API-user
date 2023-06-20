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

func getUserID(ctx context.Context) uint {
	id, ok := ctx.Value(UserIDKey).(uint)
	if !ok {
		return 0
	}
	return id
}

type userHandler struct {
	service model.UserService
}

func NewUser(us model.UserService) *userHandler {
	return &userHandler{
		service: us,
	}
}

func (u *userHandler) Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.URL.Path != "/user" {
		writeError(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		u.get(w, r)
	case http.MethodPost:
		u.post(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed)
	}
}

func (u *userHandler) User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		writeError(w, http.StatusNotFound)
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
		writeError(w, http.StatusMethodNotAllowed)
	}
}
