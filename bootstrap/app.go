package bootstrap

import (
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Application struct {
	Env *Env
	DB  *pgx.Conn
	RDB *redis.Client
}

func App() *Application {
	app := new(Application)

	logrus.Info("Init environment")
	app.Env = NewEnv()

	logrus.Info("Init database")
	app.DB = NewPostgreSQL(app.Env)

	app.RDB = NewRedis(app.Env)

	return app
}
