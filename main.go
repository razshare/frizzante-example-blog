package main

import (
	"embed"
	"github.com/razshare/frizzante/guard"
	"github.com/razshare/frizzante/route"
	"github.com/razshare/frizzante/server"
	"main/lib/routes/guards"
	"main/lib/routes/handlers/article_form"
	"main/lib/routes/handlers/board"
	"main/lib/routes/handlers/expired"
	"main/lib/routes/handlers/fallback"
	"main/lib/routes/handlers/login"
	"main/lib/routes/handlers/logout"
	"main/lib/routes/handlers/register"
)

//go:embed app/dist
var efs embed.FS
var conf = server.Default()

func main() {
	defer server.Start(conf)
	conf.Container.Efs = efs
	conf.Guards = []guard.Guard{
		{Name: "verified", Handler: guards.Verified, Tags: []string{"protected"}},
		{Name: "active", Handler: guards.Active, Tags: []string{"active"}},
	}
	conf.Routes = []route.Route{
		{Pattern: "GET /", Handler: fallback.View},
		{Pattern: "GET /expired", Handler: expired.View},
		{Pattern: "GET /login", Handler: login.View},
		{Pattern: "POST /login", Handler: login.Action},
		{Pattern: "GET /logout", Handler: logout.Action},
		{Pattern: "GET /register", Handler: register.View},
		{Pattern: "POST /register", Handler: register.Action},
		{Pattern: "GET /board", Handler: board.View},
		// Order matters here, first check for "protected", then for "active".
		// This way a session that is verified but expired, sees the message "Your sessions has expired",
		// while a session that has never been verified to begin with, is redirected to the login page.
		{Pattern: "GET /article-form", Handler: article_form.View, Tags: []string{"protected", "active"}},
		{Pattern: "POST /article-form", Handler: article_form.Action, Tags: []string{"protected", "active"}},
		{Pattern: "GET /article-remove", Handler: article_form.Remove, Tags: []string{"protected", "active"}},
	}
}
