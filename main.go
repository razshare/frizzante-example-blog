package main

import (
	"embed"
	"main/lib/core/guards"
	"main/lib/core/routes"
	"main/lib/core/servers"
	"main/lib/guards/is_logged_in"
	"main/lib/guards/is_not_expired"
	"main/lib/routes/article"
	"main/lib/routes/board"
	"main/lib/routes/expired"
	"main/lib/routes/fallback"
	"main/lib/routes/form"
	"main/lib/routes/login"
	"main/lib/routes/logout"
	"main/lib/routes/register"
)

//go:generate make clean configure
//go:generate make package
//go:generate make types
//go:embed app/dist
var efs embed.FS
var server = servers.New()

func main() {
	defer servers.Start(server)
	server.Efs = efs
	server.Routes = []routes.Route{
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
		{Pattern: "GET /form", Handler: form.View, Guards: []guards.Guard{is_logged_in.Guard, is_not_expired.Guard}},
		{Pattern: "POST /article/add", Handler: article.Add, Guards: []guards.Guard{is_logged_in.Guard, is_not_expired.Guard}},
		{Pattern: "GET /article/remove", Handler: article.Remove, Guards: []guards.Guard{is_logged_in.Guard, is_not_expired.Guard}},
	}
}
