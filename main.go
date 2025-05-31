package main

import (
	"embed"
	frz "github.com/razshare/frizzante"
	"main/lib"
	"main/lib/handlers"
	"main/lib/notifiers"
	"time"
)

//go:embed .dist/*/**
var dist embed.FS
var public []frz.Guard
var protected = []frz.Guard{
	func(c *frz.Connection, allow func()) {
		state, _ := frz.Session(c, lib.State{})

		if !state.Verified {
			c.SendNavigate("/login")
			return
		}

		allow()
	},
	func(c *frz.Connection, allow func()) {
		state, operator := frz.Session(c, lib.State{})
		defer operator.Save(state)

		if time.Since(state.LastActivity) > 30*time.Minute {
			c.SendNavigate("/expired")
			return
		}

		state.LastActivity = time.Now()
		allow()
	},
}

func main() {
	frz.NewServer().
		WithNotifier(notifiers.Console).
		WithAddress("127.0.0.1:8080").
		WithDist(dist).
		Map(public, "GET /", handlers.GetDefault).
		Map(public, "GET /expired", handlers.GetExpired).
		Map(public, "GET /login", handlers.GetLogin).
		Map(public, "POST /login", handlers.PostLogin).
		Map(public, "GET /logout", handlers.GetLogout).
		Map(public, "GET /register", handlers.GetRegister).
		Map(public, "POST /register", handlers.PostRegister).
		Map(protected, "GET /account", handlers.GetAccount).
		Map(protected, "GET /board", handlers.GetBoard).
		Map(protected, "POST /board", handlers.PostBoard).
		Start()
}
