package bootstrap

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func NewRedis(env *Env) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", env.RDBHost, env.RDBPort),
		Password: "",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		logrus.Fatalln(err)
	}
	return rdb
}
