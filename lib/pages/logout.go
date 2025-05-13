package pages

import (
	f "github.com/razshare/frizzante"
	"main/lib/guards"
)

func Logout(page *f.Page) {
	f.PageWithPath(page, "/logout")
	f.PageWithView(page, f.ViewReference("Logout"))
	f.PageWithGuardHandler(page, guards.Verified)
	f.PageWithBaseHandler(page, func(request *f.Request, response *f.Response, _ *f.View) {
		session := f.SessionStart(request, response)
		f.SessionSetBool(session, "verified", false)
		f.ResponseSendNavigate(response, "Login")
	})
	f.PageWithActionHandler(page, func(request *f.Request, response *f.Response, _ *f.View) {
		session := f.SessionStart(request, response)
		f.SessionSetBool(session, "verified", false)
		f.ResponseSendNavigate(response, "Login")
	})
}
