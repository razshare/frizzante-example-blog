package main

import (
	"embed"
	"main/lib/core/guard"
	"main/lib/core/route"
	"main/lib/core/server"
	"main/lib/core/tag"
	"main/lib/core/view/ssr"
	"main/lib/guards"
	"main/lib/routes/article"
	"main/lib/routes/board"
	"main/lib/routes/expired"
	"main/lib/routes/fallback"
	"main/lib/routes/form"
	"main/lib/routes/login"
	"main/lib/routes/logout"
	"main/lib/routes/register"
	"os"
)

const IsLoggedIn tag.Tag = 0
const IsNotExpired tag.Tag = 1

//go:embed app/dist
var efs embed.FS
var srv = server.New()
var dev = os.Getenv("DEV") == "1"
var render = ssr.New(ssr.Config{Efs: efs, UseDisk: dev})

func main() {
	defer server.Start(srv)
	srv.Efs = efs
	srv.Render = render
	srv.Guards = []guard.Guard{
		{Name: "logged-in", Handler: guards.IsLoggedIn, Tags: []tag.Tag{IsLoggedIn}},
		{Name: "active", Handler: guards.IsNotExpired, Tags: []tag.Tag{IsNotExpired}},
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
		// Order matters here, first check that the user is logged in and then that the login is nor expired.
		// This way a user that is logged in but has an expired session, sees the message "Your sessions has expired",
		// while a user that is not logged in to begin with, is redirected to the login page.
		{Pattern: "GET /form", Handler: form.View, Tags: []tag.Tag{IsLoggedIn, IsNotExpired}},
		{Pattern: "POST /form/add", Handler: form.Add, Tags: []tag.Tag{IsLoggedIn, IsNotExpired}},
		{Pattern: "GET /article/remove", Handler: article.Remove, Tags: []tag.Tag{IsLoggedIn, IsNotExpired}},
	}
}
