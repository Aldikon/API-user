package bootstrap

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Env struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`

	DBUser       string `mapstructure:"POSTGRES_USER"`
	DBPass       string `mapstructure:"POSTGRES_PASSWORD"`
	DBHost       string `mapstructure:"POSTGRES_HOST"`
	DBPort       string `mapstructure:"POSTGRES_PORT"`
	DBName       string `mapstructure:"POSTGRES_NAME"`
	DBMigPath    string `mapstructure:"POSTGRES_MIGRATION_PATH"`
	DBMigVersion int    `mapstructure:"POSTGRES_MIGRATION_VERSION"`

	RDBHost string `mapstructure:"REDIS_HOST"`
	RDBPort string `mapstructure:"REDIS_PORT"`
}

func NewEnv() *Env {
	env := new(Env)

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}

	if err := viper.Unmarshal(env); err != nil {
		logrus.Fatal(err)
	}

	return env
}
