package handlers

import "net/http"

func Index(w http.ResponseWriter, _ *http.Request) {
	writeError(w, http.StatusNotFound)
}
