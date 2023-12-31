package middleware

import (
	"net/http"
)

func Use(next http.Handler) http.Handler {
	return logger(recoverHandler(next))
}

func recoverHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := recover(); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// logrus.Tracef("path: %s\nmethod: %s\n time: %s\n", r.URL.Path, r.Method, time.Now().Format(time.DateTime))
		next.ServeHTTP(w, r)
	})
}
