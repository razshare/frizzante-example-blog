package pages

import (
	"crypto/sha256"
	"fmt"
	f "github.com/razshare/frizzante"
	"main/lib/sql"
	"time"
)

func Login(page *f.Page) {
	f.PageWithPath(page, "/")
	f.PageWithPath(page, "/login")
	f.PageWithView(page, f.ViewReference("Login"))
	f.PageWithGuardHandler(page, func(request *f.Request, response *f.Response, pass func()) {
		session := f.SessionStart(request, response)
		if f.SessionHas(session, "lastActivity") {
			lastActivity := f.SessionGetTime(session, "lastActivity")
			if time.Since(lastActivity) > 30*time.Minute {
				f.SessionDestroy(session)

				verified := f.SessionHas(session, "verified") && f.SessionGetBool(session, "verified")

				if verified {
					f.ResponseSendNavigate(response, "Expired")
				}
			}
		}
		pass()
	})

	f.PageWithActionHandler(page, func(request *f.Request, response *f.Response, view *f.View) {
		form := f.RequestReceiveForm(request)
		id := form.Get("id")
		password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

		if !sql.VerifyAccount(id, password) {
			f.ViewWithData(view, "error", "Invalid credentials")
			return
		}

		session := f.SessionStart(request, response)
		f.SessionSetBool(session, "verified", true)

		f.ResponseSendNavigate(response, "Board")
	})
}
