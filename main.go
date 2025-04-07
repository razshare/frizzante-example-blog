package main

import (
	sqlib "database/sql"
	"embed"
	_ "github.com/go-sql-driver/mysql"
	f "github.com/razshare/frizzante"
	"main/lib"
	"main/lib/guards"
	"main/lib/indexes"
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

	// Guards.
	f.ServerWithIndexGuard(s, guards.Session)
	f.ServerWithIndexGuard(s, guards.Render)

	// Indexes.
	f.ServerWithIndex(s, indexes.Board)
	f.ServerWithIndex(s, indexes.Login)
	f.ServerWithIndex(s, indexes.Logout)
	f.ServerWithIndex(s, indexes.Register)
	f.ServerWithIndex(s, indexes.Default)

	// Start.
	f.ServerStart(s)
}
