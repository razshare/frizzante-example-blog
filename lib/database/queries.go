package database

import (
	"database/sql"
	"embed"
	"log"
	"main/lib/database/sqlc"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/razshare/frizzante/files"
)

//go:embed source.sqlite
var Efs embed.FS
var Queries *sqlc.Queries

func init() {
	if !files.IsFile("source.sqlite") {
		data, readError := Efs.ReadFile("source.sqlite")
		if readError != nil {
			log.Fatal(readError)
		}
		writeError := os.WriteFile("source.sqlite", data, os.ModePerm)
		if writeError != nil {
			log.Fatal(writeError)
		}
	}

	db, dbError := sql.Open("sqlite3", "file:source.sqlite?cache=shared")
	if dbError != nil {
		log.Fatal(dbError)
	}

	Queries = sqlc.New(db)
}
