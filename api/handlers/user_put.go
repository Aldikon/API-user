package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Aldikon/API-user/model"
	"github.com/sirupsen/logrus"
)

func (u *userHandler) update(w http.ResponseWriter, r *http.Request) {
	id := getUserID(r.Context())
	if id == 0 {
		writeError(w, http.StatusNotFound)
		return
	}
	user := model.UserUpdate{
		ID: id,
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		writeError(w, http.StatusInternalServerError)
		return
	}

	err = u.service.Update(r.Context(), user)
	if err != nil {
		logrus.Trace(err)
		var e *model.ErrorValid
		if errors.As(err, &e) {
			writeError(w, http.StatusBadRequest)
			return
		}
		writeError(w, http.StatusInternalServerError)
		return
	}
}
