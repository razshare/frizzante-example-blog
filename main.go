package main

import (
	"embed"
	"github.com/razshare/frizzante/environments"
	"github.com/razshare/frizzante/guards"
	"github.com/razshare/frizzante/routes"
	"github.com/razshare/frizzante/servers"
	"github.com/razshare/frizzante/traces"
	"main/lib/handlers"
	"os"
)

//go:embed app/dist
var efs embed.FS
var server = servers.New()

func main() {
	server.Efs = efs

	if err := environments.LoadDotenv(".env"); err != nil {
		traces.Trace(server.ErrorLog, err)
	} else {
		server.Address = os.Getenv("server.address")
		server.SecureAddress = os.Getenv("server.secure_address")
		server.Key = os.Getenv("server.key")
		server.Certificate = os.Getenv("server.certificate")
		server.PublicRoot = os.Getenv("server.public_root")
		server.AppRoot = os.Getenv("server.app_root")
		server.ServerJs = os.Getenv("server.server_js")
		server.IndexHtml = os.Getenv("server.index_html")
	}

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
