package handlers

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotFound)
}
