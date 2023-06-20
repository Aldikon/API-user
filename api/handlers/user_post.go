package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Aldikon/API-user/model"
	"github.com/sirupsen/logrus"
)

func (u *userHandler) post(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Trace(err)
		writeError(w, http.StatusInternalServerError)
		return
	}

	user := model.CreateUser{}
	err = json.Unmarshal(data, &user)
	if err != nil {
		logrus.Trace(err)
		writeError(w, http.StatusBadRequest)
		return
	}

	err = u.service.Creat(r.Context(), user)
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
