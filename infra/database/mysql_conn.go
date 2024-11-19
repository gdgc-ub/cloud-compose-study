package database

import (
	"fmt"
	"log"

	"github.com/devanfer02/go-blog/infra/env"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewMySQLConn() *sqlx.DB{
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		env.AppEnv.DBUser,
		env.AppEnv.DBPass,
		env.AppEnv.DBHost,
		env.AppEnv.DBPort,
		env.AppEnv.DBName,
	)

	db, err := sqlx.Connect("mysql", dsn)

	if err != nil {
		log.Fatal("ERR: " + err.Error())
	}

	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})

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

	return db 
}