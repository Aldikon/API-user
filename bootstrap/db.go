package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/golang-migrate/migrate/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

// const driverName = "postgres"

func NewPostgreSQL(env *Env) *pgx.Conn {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	dataSourceURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		env.DBUser, env.DBPass, env.DBHost, env.DBPort, env.DBName)

	// config, err := pgx.ParseConnectionString(dataSourceName)
	// if err != nil {
	// 	logrus.Fatal(err)
	// }

	conn, err := pgx.Connect(ctx, dataSourceURL)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := conn.Ping(ctx); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Init migration")
	migrateVersion(env)

	return conn
}

func migrateVersion(env *Env) {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		env.DBUser, env.DBPass, env.DBHost, env.DBPort, env.DBName)
	pathMi := fmt.Sprintf("file://%s", env.DBMigPath)

	src, err := (&file.File{}).Open(pathMi)
	if err != nil {
		logrus.Fatalln(err)
	}

	srcDB, err := (&postgres.Postgres{}).Open(psqlInfo)
	if err != nil {
		logrus.Fatalln(err)
	}

	m, err := migrate.NewWithInstance("file", src, "postgres", srcDB)
	if err != nil {
		logrus.Fatalf("Connet db on migration: %v\n", err)
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			logrus.Fatalf("Migration up: %v\n", err)
		}
	}
	// if err := m.Force(env.DBMigVersion); err != nil {
	// 	logrus.Fatalf("Migration up: %v\n", err)
	// }
}
