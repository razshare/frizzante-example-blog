package main

import (
	"embed"
	"github.com/razshare/frizzante/frz"
	"main/lib/guards"
	"main/lib/handlers"
	"main/lib/notifiers"
)

//go:embed app/dist
var dist embed.FS

func main() {
	frz.NewServer().
		WithEfs(dist).
		WithNotifier(notifiers.Console).
		AddGuard(frz.Guard{Name: "verified", Handler: guards.Verified, Tags: []string{"protected"}}).
		AddGuard(frz.Guard{Name: "active", Handler: guards.Active, Tags: []string{"protected"}}).
		AddRoute(frz.Route{Pattern: "GET /", Handler: handlers.Default}).
		AddRoute(frz.Route{Pattern: "GET /expired", Handler: handlers.Expired}).
		AddRoute(frz.Route{Pattern: "GET /login", Handler: handlers.Login}).
		AddRoute(frz.Route{Pattern: "POST /login", Handler: handlers.LoginAction}).
		AddRoute(frz.Route{Pattern: "GET /logout", Handler: handlers.LogoutAction}).
		AddRoute(frz.Route{Pattern: "GET /register", Handler: handlers.Register}).
		AddRoute(frz.Route{Pattern: "POST /register", Handler: handlers.RegisterAction}).
		AddRoute(frz.Route{Pattern: "GET /board", Handler: handlers.Board}).
		AddRoute(frz.Route{Pattern: "GET /article-form", Handler: handlers.ArticleForm, Tags: []string{"protected"}}).
		AddRoute(frz.Route{Pattern: "POST /article-form", Handler: handlers.ArticleFormAction, Tags: []string{"protected"}}).
		Start()
}
