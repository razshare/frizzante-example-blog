package main

import (
	"embed"
	"github.com/razshare/frizzante/web"
	"main/lib"
	"main/lib/handlers"
)

//go:embed app/dist
var efs embed.FS
var server = web.NewServer()

func main() {
	server.Efs = efs
	server.Notifier = lib.Notifier
	server.AddGuard(web.Guard{Name: "verified", Handler: handlers.GuardVerified, Tags: []string{"protected"}})
	server.AddGuard(web.Guard{Name: "active", Handler: handlers.GuardActive, Tags: []string{"active"}})
	server.AddRoute(web.Route{Pattern: "GET /", Handler: handlers.Default})
	server.AddRoute(web.Route{Pattern: "GET /expired", Handler: handlers.Expired})
	server.AddRoute(web.Route{Pattern: "GET /login", Handler: handlers.Login})
	server.AddRoute(web.Route{Pattern: "POST /login", Handler: handlers.LoginAction})
	server.AddRoute(web.Route{Pattern: "GET /logout", Handler: handlers.LogoutAction})
	server.AddRoute(web.Route{Pattern: "GET /register", Handler: handlers.Register})
	server.AddRoute(web.Route{Pattern: "POST /register", Handler: handlers.RegisterAction})
	server.AddRoute(web.Route{Pattern: "GET /board", Handler: handlers.Board})

	// Order matters here, first check for "protected", then for "active".
	// This way a sessions that is verified but expired, sees the message "Your sessions has expired",
	// while a sessions that has never been verified to begin with, is redirected to the login page.
	server.AddRoute(web.Route{Pattern: "GET /article-form", Handler: handlers.ArticleForm, Tags: []string{"protected", "active"}})
	server.AddRoute(web.Route{Pattern: "POST /article-form", Handler: handlers.ArticleFormAction, Tags: []string{"protected", "active"}})
	server.Start()
}
