package main

import (
	sqlib "database/sql"
	"embed"
	_ "github.com/go-sql-driver/mysql"
	f "github.com/razshare/frizzante"
	"main/lib"
	"main/lib/pages"
	"main/lib/sql"
)

//go:embed .dist/*/**
var d embed.FS

func main() {
	db, dbe := sqlib.Open("mysql", "root:root@/forum")
	if dbe != nil {
		panic(dbe)
	}

	// Sql.
	f.SqlWithDatabase(sql.Sql, db)
	f.SqlWithNotifier(sql.Sql, lib.Notifier)

	// Server.
	s := f.ServerCreate()
	f.ServerWithPort(s, 8080)
	f.ServerWithHostName(s, "127.0.0.1")
	f.ServerWithEmbeddedFileSystem(s, d)

	// Pages.
	f.ServerWithPageBuilder(s, pages.Board)
	f.ServerWithPageBuilder(s, pages.Login)
	f.ServerWithPageBuilder(s, pages.Logout)
	f.ServerWithPageBuilder(s, pages.Register)
	f.ServerWithPageBuilder(s, pages.Default)

	// Start.
	f.ServerStart(s)
}
