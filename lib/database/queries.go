package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"main/lib/utilities/sqlc"
)

var Queries *sqlc.Queries

func init() {
	db, dbError := sql.Open("sqlite3", "file:database.sqlite?cache=shared")
	if dbError != nil {
		log.Fatal(dbError)
	}

	Queries = sqlc.New(db)
}
