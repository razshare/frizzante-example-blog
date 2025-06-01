package main

import (
	"embed"
	frz "github.com/razshare/frizzante"
	"main/lib/guards"
	"main/lib/handlers"
	"main/lib/notifiers"
)

//go:embed .dist/*/**
var dist embed.FS

func main() {
	frz.NewServer().
		WithDist(dist).
		WithAddress("127.0.0.1:8080").
		WithNotifier(notifiers.Console).
		AddGuard(frz.Guard{Handler: guards.Verified, Tags: []string{"protected"}}).
		AddGuard(frz.Guard{Handler: guards.Active, Tags: []string{"protected"}}).
		AddRoute(frz.Route{Pattern: "GET /", Handler: handlers.GetDefault}).
		AddRoute(frz.Route{Pattern: "GET /expired", Handler: handlers.GetExpired}).
		AddRoute(frz.Route{Pattern: "GET /login", Handler: handlers.GetLogin}).
		AddRoute(frz.Route{Pattern: "POST /login", Handler: handlers.PostLogin}).
		AddRoute(frz.Route{Pattern: "GET /logout", Handler: handlers.GetLogout}).
		AddRoute(frz.Route{Pattern: "GET /register", Handler: handlers.GetRegister}).
		AddRoute(frz.Route{Pattern: "POST /register", Handler: handlers.PostRegister}).
		AddRoute(frz.Route{Pattern: "GET /account", Handler: handlers.GetAccount, Tags: []string{"protected"}}).
		AddRoute(frz.Route{Pattern: "GET /board", Handler: handlers.GetBoard, Tags: []string{"protected"}}).
		AddRoute(frz.Route{Pattern: "POST /board", Handler: handlers.PostBoard, Tags: []string{"protected"}}).
		Start()
}
