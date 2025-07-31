package main

import (
	"embed"
	"github.com/razshare/frizzante/guards"
	"github.com/razshare/frizzante/routes"
	"github.com/razshare/frizzante/servers"
	"main/lib/handlers"
)

//go:embed app/dist
var efs embed.FS
var server = servers.New()

func main() {
	server.Efs = efs

	server.Guards = []guards.Guard{
		{Name: "verified", Handler: handlers.GuardVerified, Tags: []string{"protected"}},
		{Name: "active", Handler: handlers.GuardActive, Tags: []string{"active"}},
	}

	server.Routes = []routes.Route{
		{Pattern: "GET /", Handler: handlers.Default},
		{Pattern: "GET /expired", Handler: handlers.Expired},
		{Pattern: "GET /login", Handler: handlers.Login},
		{Pattern: "POST /login", Handler: handlers.LoginAction},
		{Pattern: "GET /logout", Handler: handlers.LogoutAction},
		{Pattern: "GET /register", Handler: handlers.Register},
		{Pattern: "POST /register", Handler: handlers.RegisterAction},
		{Pattern: "GET /board", Handler: handlers.Board},

		// Order matters here, first check for "protected", then for "active".
		// This way a session that is verified but expired, sees the message "Your sessions has expired",
		// while a session that has never been verified to begin with, is redirected to the login page.
		{Pattern: "GET /article-form", Handler: handlers.ArticleForm, Tags: []string{"protected", "active"}},
		{Pattern: "POST /article-form", Handler: handlers.ArticleFormAction, Tags: []string{"protected", "active"}},
		{Pattern: "GET /article-remove", Handler: handlers.ArticleRemove, Tags: []string{"protected", "active"}},
	}

	server.Start()
}
