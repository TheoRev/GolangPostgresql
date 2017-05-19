package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// obtiene la conexion la base de datos
func getConnection() *sql.DB {
	dsn := "postgres://theo:ambu@localhost:5432/gocrud?sslmode=disable"
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}
	return db
}
