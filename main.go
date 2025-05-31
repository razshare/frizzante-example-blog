package main

import (
	"embed"
	"github.com/razshare/frizzante"
	"main/lib/handlers"
	"main/lib/notifiers"
)

//go:embed .dist/*/**
var dist embed.FS
var server = frizzante.
	NewServer().
	WithNotifier(notifiers.Console).
	WithAddress("127.0.0.1:8080")

func main() {
	server.
		WithDist(dist).
		WithRequestHandler("GET /", handlers.GetDefault).
		WithRequestHandler("GET /account", handlers.GetAccount).
		WithRequestHandler("GET /board", handlers.GetBoard).
		WithRequestHandler("POST /board", handlers.PostBoard).
		WithRequestHandler("GET /expired", handlers.GetExpired).
		WithRequestHandler("GET /login", handlers.GetLogin).
		WithRequestHandler("POST /login", handlers.PostLogin).
		WithRequestHandler("GET /logout", handlers.GetLogout).
		WithRequestHandler("GET /register", handlers.GetRegister).
		WithRequestHandler("POST /register", handlers.PostRegister).
		Start()
}
