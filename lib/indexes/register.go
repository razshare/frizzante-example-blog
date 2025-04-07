package indexes

import (
	"crypto/sha256"
	"fmt"
	f "github.com/razshare/frizzante"
	"main/lib/sql"
)

func registerShowFunction(_ *f.Request, _ *f.Response, _ *f.Page) {
	// Noop.
}

func registerActionFunction(req *f.Request, res *f.Response, _ *f.Page) {
	form := f.ReceiveForm(req)
	id := form.Get("id")
	dn := form.Get("displayName")
	pwd := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

	sql.AddAccount(id, dn, pwd)
	f.SendNavigate(res, "login")
}

func Register(
	route func(path string, page string),
	show func(showFunction f.PageFunction),
	action func(actionFunction f.PageFunction),
) {
	route("/register", "register")
	show(registerShowFunction)
	action(registerActionFunction)
}
