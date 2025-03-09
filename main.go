package main

import (
	sqlib "database/sql"
	"embed"
	_ "github.com/go-sql-driver/mysql"
	frz "github.com/razshare/frizzante"
	"log"
	"main/pages"
	"main/schemas"
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
	frz.SqlWithDatabase(schemas.Sql, db)
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
	frz.SqlRecallError(schemas.Sql, func(err error) {
		lg.Fatal(err)
	})

	// Drop tables.
	frz.SqlDropTable[schemas.CommentContent](schemas.Sql)
	frz.SqlDropTable[schemas.ArticleContent](schemas.Sql)
	frz.SqlDropTable[schemas.Comment](schemas.Sql)
	frz.SqlDropTable[schemas.Article](schemas.Sql)
	frz.SqlDropTable[schemas.Account](schemas.Sql)

	// Create tables.
	frz.SqlCreateTable[schemas.Account](schemas.Sql)
	frz.SqlCreateTable[schemas.Article](schemas.Sql)
	frz.SqlCreateTable[schemas.Comment](schemas.Sql)
	frz.SqlCreateTable[schemas.ArticleContent](schemas.Sql)
	frz.SqlCreateTable[schemas.CommentContent](schemas.Sql)

	// Start.
	frz.ServerStart(srv)
}
