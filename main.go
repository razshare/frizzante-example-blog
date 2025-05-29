package main

import (
	"embed"
	"github.com/razshare/frizzante"
	"main/lib/notifiers"
	"main/lib/routes"
)

//go:embed .dist/*/**
var dist embed.FS
var server = frizzante.
	NewServer().
	WithNotifier(notifiers.Console).
	WithAddress("127.0.0.1:8080")

func main() {
	server.
		WithRequestHandler("GET /", routes.GetDefault).
		WithRequestHandler("GET /account", routes.GetAccount).
		WithRequestHandler("GET /board", routes.GetBoard).
		WithRequestHandler("POST /board", routes.PostBoard).
		WithRequestHandler("GET /expired", routes.GetExpired).
		WithRequestHandler("GET /login", routes.GetLogin).
		WithRequestHandler("POST /login", routes.PostLogin).
		WithRequestHandler("GET /logout", routes.GetLogout).
		WithRequestHandler("GET /register", routes.GetRegister).
		WithRequestHandler("POST /register", routes.PostRegister).
		Start(dist)
}
