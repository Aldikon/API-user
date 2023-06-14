package main

import (
	"github.com/Aldikon/API-user/api/server"
	"github.com/Aldikon/API-user/bootstrap"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func main() {
	logrus.Info("Start work program")
	app := bootstrap.App()

	logrus.Info("Init server")
	server := server.New(app)

	logrus.Infof("Start listen http://localhost%v", app.Env.ServerAddress)
	server.ListenAndServe()
}
