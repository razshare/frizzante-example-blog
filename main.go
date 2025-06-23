package main

import (
	"embed"
	"github.com/razshare/frizzante/libsrv"
	"main/lib/guards"
	"main/lib/handlers"
	"main/lib/notifiers"
)

//go:embed app/dist
var efs embed.FS
var server = libsrv.NewServer()

func main() {
	server.Efs = efs
	server.Notifier = notifiers.Console
	server.AddGuard(libsrv.Guard{Name: "verified", Handler: guards.Verified, Tags: []string{"protected"}})
	server.AddGuard(libsrv.Guard{Name: "active", Handler: guards.Active, Tags: []string{"protected", "active"}})
	server.AddRoute(libsrv.Route{Pattern: "GET /", Handler: handlers.Default})
	server.AddRoute(libsrv.Route{Pattern: "GET /expired", Handler: handlers.Expired})
	server.AddRoute(libsrv.Route{Pattern: "GET /login", Handler: handlers.Login})
	server.AddRoute(libsrv.Route{Pattern: "POST /login", Handler: handlers.LoginAction})
	server.AddRoute(libsrv.Route{Pattern: "GET /logout", Handler: handlers.LogoutAction})
	server.AddRoute(libsrv.Route{Pattern: "GET /register", Handler: handlers.Register})
	server.AddRoute(libsrv.Route{Pattern: "POST /register", Handler: handlers.RegisterAction})
	server.AddRoute(libsrv.Route{Pattern: "GET /board", Handler: handlers.Board, Tags: []string{"active"}})
	server.AddRoute(libsrv.Route{Pattern: "GET /article-form", Handler: handlers.ArticleForm, Tags: []string{"protected"}})
	server.AddRoute(libsrv.Route{Pattern: "POST /article-form", Handler: handlers.ArticleFormAction, Tags: []string{"protected"}})
	server.Start()
}
