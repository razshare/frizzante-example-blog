package main

import (
	sqlib "database/sql"
	"embed"
	_ "github.com/go-sql-driver/mysql"
	frz "github.com/razshare/frizzante"
	"main/lib"
	"main/lib/pages"
)

//go:embed .dist/*/**
var dist embed.FS

func main() {
	database, databaseError := sqlib.Open("mysql", "root:root@/forum")
	if databaseError != nil {
		panic(databaseError)
	}

	// Sql.
	frz.SqlWithDatabase(lib.Sql, database)
	frz.SqlWithNotifier(lib.Sql, lib.Notifier)

	// Server.
	server := frz.ServerCreate()
	frz.ServerWithPort(server, 8080)
	frz.ServerWithHostName(server, "127.0.0.1")
	frz.ServerWithEmbeddedFileSystem(server, dist)

	// Route (order matters, "/" should always be last).
	frz.ServerRoutePage(server, "GET /register", "Register", pages.Register)
	frz.ServerRoutePage(server, "POST /register", "Register", pages.Register)
	frz.ServerRoutePage(server, "GET /login", "Login", pages.Login)
	frz.ServerRoutePage(server, "POST /login", "Login", pages.Login)
	frz.ServerRoutePage(server, "GET /", "Login", pages.Login)
	frz.ServerRoutePage(server, "POST /", "Login", pages.Login)

	// Start.
	frz.ServerStart(server)
}
