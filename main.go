package main

import (
	"embed"
	"github.com/razshare/frizzante/guards"
	"github.com/razshare/frizzante/routes"
	"github.com/razshare/frizzante/server"
	"main/lib"
	"main/lib/handlers"
)

//go:embed app/dist
var efs embed.FS

func main() {
	server.WithEfs(efs)
	server.WithNotifier(lib.Notifier)
	server.AddGuard(guards.Guard{Name: "verified", Handler: handlers.GuardVerified, Tags: []string{"protected"}})
	server.AddGuard(guards.Guard{Name: "active", Handler: handlers.GuardActive, Tags: []string{"active"}})
	server.AddRoute(routes.Route{Pattern: "GET /", Handler: handlers.Default})
	server.AddRoute(routes.Route{Pattern: "GET /expired", Handler: handlers.Expired})
	server.AddRoute(routes.Route{Pattern: "GET /login", Handler: handlers.Login})
	server.AddRoute(routes.Route{Pattern: "POST /login", Handler: handlers.LoginAction})
	server.AddRoute(routes.Route{Pattern: "GET /logout", Handler: handlers.LogoutAction})
	server.AddRoute(routes.Route{Pattern: "GET /register", Handler: handlers.Register})
	server.AddRoute(routes.Route{Pattern: "POST /register", Handler: handlers.RegisterAction})
	server.AddRoute(routes.Route{Pattern: "GET /board", Handler: handlers.Board})

	// Order matters here, first check for "protected", then for "active".
	// This way a sessions that is verified but expired, sees the message "Your sessions has expired",
	// while a sessions that has never been verified to begin with, is redirected to the login page.
	server.AddRoute(routes.Route{Pattern: "GET /article-form", Handler: handlers.ArticleForm, Tags: []string{"protected", "active"}})
	server.AddRoute(routes.Route{Pattern: "POST /article-form", Handler: handlers.ArticleFormAction, Tags: []string{"protected", "active"}})
	server.Start()
}
