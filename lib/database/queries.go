package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"main/lib/generated"
	"main/lib/value"
)

var Queries = generated.New(value.WrapFatal(sql.Open("sqlite3", "file:database.sqlite?cache=shared")).Value)
