package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (u *userHandler) delete(w http.ResponseWriter, r *http.Request) {
	id := getUserID(r.Context())
	if id == 0 {
		logrus.Trace("id null")
		writeError(w, http.StatusNotFound)
		return
	}

	err := u.service.Delete(r.Context(), id)
	if err != nil {
		logrus.Trace(err)
		writeError(w, http.StatusInternalServerError)
		return
	}
}
