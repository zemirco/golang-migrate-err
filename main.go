package main

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// use with local kubernetes cluster on minikube
	// source := "user=postgres host=192.168.99.100 port=32283 dbname=postgres sslmode=disable"
	// use with Circle CI
	source := "user=postgres host=localhost dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", source)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	migration, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		panic(err)
	}
	if err := migration.Drop(); err != nil {
		panic(err)
	}
	if err := migration.Up(); err != nil {
		panic(err)
	}
}
