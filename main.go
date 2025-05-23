package main

import (
	"embed"
	_ "github.com/go-sql-driver/mysql"
	"main/lib/config"
	_ "main/lib/config"
	_ "main/lib/controllers/account"
	_ "main/lib/controllers/any"
	_ "main/lib/controllers/board"
	_ "main/lib/controllers/expired"
	_ "main/lib/controllers/login"
	_ "main/lib/controllers/logout"
	_ "main/lib/controllers/register"
)

//go:embed .dist/*/**
var dist embed.FS

func main() {
	config.Server.Start(dist)
}
