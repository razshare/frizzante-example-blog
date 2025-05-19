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
	// Errors.
	if databaseError != nil {
		log.Fatal(databaseError)
	}

	// Create.
	server := f.NewServer()
	notifier := f.NewNotifier()

	// Configure.
	server.WithPort(8080)
	server.WithNotifier(notifier)
	server.WithHostName("127.0.0.1")
	server.WithEmbeddedFileSystem(&dist)

	// Sql.
	lib.Sql.WithNotifier(notifier)
	lib.Sql.WithDatabase(database)

	// Pages.
	server.WithPageController(pages.BoardController{})
	server.WithPageController(pages.ExpiredController{})
	server.WithPageController(pages.LoginController{})
	server.WithPageController(pages.LogoutController{})
	server.WithPageController(pages.RegisterController{})
	server.WithPageController(pages.DefaultController{})
	server.WithPageController(pages.AccountController{})

	//Start.
	server.Start()
}
