package pages

import (
	"crypto/sha256"
	"fmt"
	f "github.com/razshare/frizzante"
	"main/lib/sql"
)

func Login(context f.PageContext) {
	// Context.
	path, view, base, action := context()

	// Configure.
	path("/login")
	view(f.ViewReference("Login"))
	base(func(_ *f.Request, _ *f.Response, _ *f.View) {
		// Noop.
	})
	action(func(request *f.Request, response *f.Response, view *f.View) {
		_, set, _ := f.SessionStart(request, response)

		form := f.RequestReceiveForm(request)
		id := form.Get("id")
		password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

		if !sql.VerifyAccount(id, password) {
			f.ViewWithData(view, "error", "Invalid credentials")
			return
		}
		set("verified", true)

		f.ResponseSendNavigate(response, "Board")
	})
}
