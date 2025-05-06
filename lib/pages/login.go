package pages

import (
	"crypto/sha256"
	"fmt"
	f "github.com/razshare/frizzante"
	"main/lib/sql"
)

func loginBaseFunction(_ *f.Request, _ *f.Response, _ *f.View) {
	// Noop.
}

func loginActionFunction(request *f.Request, response *f.Response, view *f.View) {
	_, set, _ := f.SessionStart(request, response)

	form := f.ReceiveForm(request)
	id := form.Get("id")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

	if !sql.VerifyAccount(id, password) {
		f.ViewWithData(view, "error", "Invalid credentials")
		return
	}
	set("verified", true)

	f.SendNavigate(response, "Board")
}

func Login(context f.PageContext) {
	path, view, base, action := context()
	path("/login")
	view(f.ViewReference("Login"))
	base(loginBaseFunction)
	action(loginActionFunction)
}
