package database

import (
	"database/sql"
	"embed"
	_ "github.com/mattn/go-sqlite3"
	"github.com/razshare/frizzante/libfs"
	"log"
	"main/lib/utilities/sqlc"
	"os"
)

//go:embed database.sqlite
var dbf embed.FS

var Queries *sqlc.Queries

func init() {
	if !libfs.IsFile("database.sqlite") {
		data, readError := dbf.ReadFile("database.sqlite")
		if readError != nil {
			log.Fatal(readError)
		}
		writeError := os.WriteFile("database.sqlite", data, os.ModePerm)
		if writeError != nil {
			log.Fatal(writeError)
		}
	}

	db, dbError := sql.Open("sqlite3", "file:database.sqlite?cache=shared")
	if dbError != nil {
		log.Fatal(dbError)
	}

	Queries = sqlc.New(db)
}
