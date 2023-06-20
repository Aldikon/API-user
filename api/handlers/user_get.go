package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Aldikon/API-user/model"
	"github.com/sirupsen/logrus"
)

func (u *userHandler) get(w http.ResponseWriter, r *http.Request) {
	filter := model.FilterUser{}

	filter.Parse(r.URL.Query())

	users, err := u.service.ListWithFilter(r.Context(), filter)
	if err != nil {
		logrus.Trace(err)
		writeError(w, http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(users)
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

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
