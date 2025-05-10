package main

import (
	sqlib "database/sql"
	"embed"
	_ "github.com/go-sql-driver/mysql"
	f "github.com/razshare/frizzante"
	"log"
	"main/lib"
	"main/lib/pages"
	"main/lib/sql"
)

//go:embed .dist/*/**
var d embed.FS

func main() {
	// Database.
	database, databaseError := sqlib.Open("mysql", "root:root@/forum")
	if databaseError != nil {
		log.Fatal(databaseError)
	}

	// Sql.
	f.SqlWithDatabase(sql.Sql, database)
	f.SqlWithNotifier(sql.Sql, lib.Notifier)

	// Server.
	server := f.ServerCreate()
	f.ServerWithPort(server, 8080)
	f.ServerWithHostName(server, "127.0.0.1")
	f.ServerWithEmbeddedFileSystem(server, d)

	// Pages.
	f.ServerWithPageBuilder(server, pages.Board)
	f.ServerWithPageBuilder(server, pages.Login)
	f.ServerWithPageBuilder(server, pages.Logout)
	f.ServerWithPageBuilder(server, pages.Register)
	f.ServerWithPageBuilder(server, pages.Default)

	// Start.
	f.ServerStart(server)
}
