package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/devanfer02/go-blog/infra/env"
)

func NewPgsqlConn() *sqlx.DB {
	dbx, err := sqlx.Connect("postgres", fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s sslmode=disable port=%s",
		env.AppEnv.DBUser,
		env.AppEnv.DBPass,
		env.AppEnv.DBHost,
		env.AppEnv.DBName,
		env.AppEnv.DBPort,
	))

	if err != nil {
		log.Fatal("ERR: " + err.Error())
	}

	driver, err := postgres.WithInstance(dbx.DB, &postgres.Config{})

	if err != nil {
		log.Fatal("ERR: " + err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./data/db/migrations",
		env.AppEnv.DBName, driver, 
	)

	if err != nil {
		log.Fatal("ERR: " + err.Error())
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("ERR: " + err.Error())
	}

	return dbx
}
