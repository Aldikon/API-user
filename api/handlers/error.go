package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type error struct {
	Status  int    `json:"status"`
	Messege string `json:"message"`
}

func writeError(w http.ResponseWriter, code int) {
	data, err := json.Marshal(error{
		Status:  code,
		Messege: http.StatusText(code),
	})
	if err != nil {
		logrus.Trace(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
