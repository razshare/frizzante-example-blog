package main

import (
	sqlib "database/sql"
	"embed"
	_ "github.com/go-sql-driver/mysql"
	frz "github.com/razshare/frizzante"
	"log"
	"main/database"
	"main/pages"
	"os"
)

//go:embed .dist/*/**
var efs embed.FS

func main() {
	// Create.
	db, err := sqlib.Open("mysql", "root:root@/forum")
	if err != nil {
		panic(err)
	}
	lg := log.Default()
	srv := frz.ServerCreate()

	// Setup.
	frz.SqlWithDatabase(database.Sql, db)
	frz.ServerWithPort(srv, 8080)
	frz.ServerWithHostName(srv, "127.0.0.1")
	frz.ServerWithEmbeddedFileSystem(srv, efs)
	frz.ServerWithLogger(srv, lg)

	// Route (order matters, "/" should always be last).
	frz.ServerRoutePage(srv, "GET /register", "register", pages.Register)
	frz.ServerRoutePage(srv, "POST /register", "register", pages.Register)
	frz.ServerRoutePage(srv, "GET /login", "login", pages.Login)
	frz.ServerRoutePage(srv, "POST /login", "login", pages.Login)
	frz.ServerRoutePage(srv, "GET /", "login", pages.Login)
	frz.ServerRoutePage(srv, "POST /", "login", pages.Login)

	// Log.
	frz.ServerRecallError(srv, func(err error) {
		lg.Fatal(err)
	})
	frz.SqlRecallError(database.Sql, func(err error) {
		lg.Fatal(err)
	})

	// Schemas.
	if "1" != os.Getenv("PROD") {
		frz.SqlDropTable[CommentContent](database.Sql)
		frz.SqlDropTable[ArticleContent](database.Sql)
		frz.SqlDropTable[Comment](database.Sql)
		frz.SqlDropTable[Article](database.Sql)
		frz.SqlDropTable[Account](database.Sql)
	}
	frz.SqlCreateTable[Account](database.Sql)
	frz.SqlCreateTable[Article](database.Sql)
	frz.SqlCreateTable[Comment](database.Sql)
	frz.SqlCreateTable[ArticleContent](database.Sql)
	frz.SqlCreateTable[CommentContent](database.Sql)

	// Start.
	frz.ServerStart(srv)
}
