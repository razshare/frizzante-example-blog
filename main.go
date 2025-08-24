package main

import (
	"embed"
	"github.com/razshare/frizzante/guard"
	"github.com/razshare/frizzante/route"
	"github.com/razshare/frizzante/server"
	"github.com/razshare/frizzante/svelte/ssr"
	"github.com/razshare/frizzante/tag"
	"main/lib/guards"
	"main/lib/routes/article"
	"main/lib/routes/article_form"
	"main/lib/routes/board"
	"main/lib/routes/expired"
	"main/lib/routes/fallback"
	"main/lib/routes/login"
	"main/lib/routes/logout"
	"main/lib/routes/register"
	"os"
)

const Active tag.Tag = 0
const Verified tag.Tag = 1

//go:embed app/dist
var efs embed.FS
var srv = server.New()
var dev = os.Getenv("DEV") == "1"
var render = ssr.New(ssr.Config{Efs: efs, Disk: dev})

func main() {
	defer server.Start(srv)
	srv.Efs = efs
	srv.Render = render
	srv.Guards = []guard.Guard{
		{Name: "verified", Handler: guards.Verified, Tags: []tag.Tag{Verified}},
		{Name: "active", Handler: guards.Active, Tags: []tag.Tag{Active}},
	}
	srv.Routes = []route.Route{
		{Pattern: "GET /", Handler: fallback.View},
		{Pattern: "GET /expired", Handler: expired.View},
		{Pattern: "GET /login", Handler: login.View},
		{Pattern: "POST /login", Handler: login.Action},
		{Pattern: "GET /logout", Handler: logout.Action},
		{Pattern: "GET /register", Handler: register.View},
		{Pattern: "POST /register", Handler: register.Action},
		{Pattern: "GET /board", Handler: board.View},
		// Order matters here, first check for "verified", then for "active".
		// This way a session that is verified but expired, sees the message "Your sessions has expired",
		// while a session that has never been verified to begin with, is redirected to the login page.
		{Pattern: "GET /article-form", Handler: article_form.View, Tags: []tag.Tag{Verified, Active}},
		{Pattern: "POST /article-form", Handler: article.Add, Tags: []tag.Tag{Verified, Active}},
		{Pattern: "GET /article-remove", Handler: article.Remove, Tags: []tag.Tag{Verified, Active}},
	}
}
