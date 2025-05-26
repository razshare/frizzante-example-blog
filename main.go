package main

import (
	"embed"
	"github.com/razshare/frizzante"
	"main/lib"
	"main/lib/routes"
)

//go:embed .dist/*/**
var dist embed.FS

func main() {
	frizzante.
		NewServer().
		WithNotifier(lib.Notifier).
		WithAddress("127.0.0.1:8080").
		WithRequestHandler("GET /account", routes.GetAccount).
		WithRequestHandler("GET /expired", routes.GetExpired).
		WithRequestHandler("GET /logout", routes.GetLogout).
		WithRequestHandler("GET /login", routes.GetLogin).
		WithRequestHandler("POST /login", routes.PostLogin).
		WithRequestHandler("GET /register", routes.GetRegister).
		WithRequestHandler("POST /register", routes.PostRegister).
		WithRequestHandler("GET /board", routes.GetBoard).
		WithRequestHandler("POST /board", routes.PostBoard).
		WithRequestHandler("GET /", func(req *frizzante.Request, res *frizzante.Response) {
			res.SendFileOrElse(func() {
				routes.GetLogin(req, res)
			})
		}).
		Start(dist)
}
