package pages

import (
	f "github.com/razshare/frizzante"
	"main/lib/guards"
)

func Default(page *f.Page) {
	f.PageWithPath(page, "/")
	f.PageWithGuardHandler(page, guards.Session)
	f.PageWithBaseHandler(page, func(_ *f.Request, response *f.Response, _ *f.View) {
		f.ResponseSendNavigate(response, "Login")
	})
	f.PageWithActionHandler(page, func(_ *f.Request, response *f.Response, _ *f.View) {
		f.ResponseSendNavigate(response, "Login")
	})
}
