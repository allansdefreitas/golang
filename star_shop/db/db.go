package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectToDataBase() *sql.DB {

	connection := "user=postgres dbname=star_shop password=postgres host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
