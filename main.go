package main

import (
	sqlib "database/sql"
	"embed"
	_ "github.com/go-sql-driver/mysql"
	frz "github.com/razshare/frizzante"
	"main/lib"
	"main/lib/guards"
	"main/lib/indexes"
	"main/lib/sql"
)

//go:embed .dist/*/**
var dist embed.FS

func main() {
	database, databaseError := sqlib.Open("mysql", "root:root@/forum")
	if databaseError != nil {
		panic(databaseError)
	}

	// Sql.
	frz.SqlWithDatabase(sql.Sql, database)
	frz.SqlWithNotifier(sql.Sql, lib.Notifier)

	// Server.
	server := frz.ServerCreate()
	frz.ServerWithPort(server, 8080)
	frz.ServerWithHostName(server, "127.0.0.1")
	frz.ServerWithEmbeddedFileSystem(server, dist)

	// Guards.
	frz.ServerWithPageGuard(server, guards.Session)
	frz.ServerWithPageGuard(server, guards.Render)

	// Routes.
	frz.ServerWithPage(server, "/board/{user}", "board", indexes.Board)
	frz.ServerWithPage(server, "/login", "login", indexes.Login)
	frz.ServerWithPage(server, "/logout", "logout", indexes.Logout)
	frz.ServerWithPage(server, "/register", "register", indexes.Register)
	frz.ServerWithPage(server, "/", "login", indexes.Login)

	// Start.
	frz.ServerStart(server)
}
