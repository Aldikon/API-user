package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (u *userHandler) getByID(w http.ResponseWriter, r *http.Request) {
	id := getUserID(r.Context())
	if id == 0 {
		logrus.Trace("id null")
		writeError(w, http.StatusInternalServerError)
		return
	}

	user, err := u.service.GetByID(r.Context(), id)
	if err != nil {
		logrus.Trace(err)
		writeError(w, http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		logrus.Trace(err)
		writeError(w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
