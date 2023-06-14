package server

import (
	"net/http"

	"github.com/Aldikon/API-user/api/middleware"
	"github.com/Aldikon/API-user/api/route"
	"github.com/Aldikon/API-user/bootstrap"
	"github.com/sirupsen/logrus"
)

type server struct {
	*http.ServeMux
}

func New(app *bootstrap.Application) *http.Server {
	mux := http.NewServeMux()

	logrus.Info("Init route")
	route.Setup(app, mux)

	return &http.Server{
		Addr:    app.Env.ServerAddress,
		Handler: middleware.Use(mux),
	}
}
