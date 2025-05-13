package pages

import (
	"crypto/sha256"
	"fmt"
	f "github.com/razshare/frizzante"
	"main/lib/sql"
)

func Login(page *f.Page) {
	f.PageWithPath(page, "/")
	f.PageWithPath(page, "/login")
	f.PageWithView(page, f.ViewReference("Login"))
	f.PageWithGuardHandler(page, func(request *f.Request, response *f.Response, pass func()) {
		f.SessionStart(request, response)
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
