package main

import (
	sqlib "database/sql"
	"embed"
	_ "github.com/go-sql-driver/mysql"
	f "github.com/razshare/frizzante"
	"log"
	"main/lib"
	"main/lib/controllers/pages"
)

//go:embed .dist/*/**
var dist embed.FS

var database, databaseError = sqlib.Open("mysql", "root:root@/forum")

func main() {
	if databaseError != nil {
		log.Fatal(databaseError)
	}
	lib.Sql.WithDatabase(database)

	f.
		NewServer().
		WithAddress("127.0.0.1:8080").
		WithEfs(dist).
		WithPageController(pages.Board).
		WithPageController(pages.Expired).
		WithPageController(pages.Login).
		WithPageController(pages.Logout).
		WithPageController(pages.Register).
		WithPageController(pages.Account).
		WithPageController(pages.Redirect).
		WithPageController(pages.Any).
		Start()
}
