package main

import (
	"embed"
	"github.com/razshare/frizzante/guards"
	"github.com/razshare/frizzante/routes"
	"github.com/razshare/frizzante/servers"
	handlers2 "main/lib/routes/handlers"
)

//go:embed app/dist
var efs embed.FS
var server = servers.New()

func main() {
	server.Efs = efs

	server.Guards = []guards.Guard{
		{Name: "verified", Handler: handlers2.GuardVerified, Tags: []string{"protected"}},
		{Name: "active", Handler: handlers2.GuardActive, Tags: []string{"active"}},
	}

	server.Routes = []routes.Route{
		{Pattern: "GET /", Handler: handlers2.Default},
		{Pattern: "GET /expired", Handler: handlers2.Expired},
		{Pattern: "GET /login", Handler: handlers2.Login},
		{Pattern: "POST /login", Handler: handlers2.LoginAction},
		{Pattern: "GET /logout", Handler: handlers2.LogoutAction},
		{Pattern: "GET /register", Handler: handlers2.Register},
		{Pattern: "POST /register", Handler: handlers2.RegisterAction},
		{Pattern: "GET /board", Handler: handlers2.Board},

		// Order matters here, first check for "protected", then for "active".
		// This way a session that is verified but expired, sees the message "Your sessions has expired",
		// while a session that has never been verified to begin with, is redirected to the login page.
		{Pattern: "GET /article-form", Handler: handlers2.ArticleForm, Tags: []string{"protected", "active"}},
		{Pattern: "POST /article-form", Handler: handlers2.ArticleFormAction, Tags: []string{"protected", "active"}},
		{Pattern: "GET /article-remove", Handler: handlers2.ArticleRemove, Tags: []string{"protected", "active"}},
	}

	server.Start()
}
