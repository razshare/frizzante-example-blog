package indexes

import (
	"crypto/sha256"
	"fmt"
	f "github.com/razshare/frizzante"
	"main/lib/sql"
)

func registerAction(req *f.Request, res *f.Response, _ *f.Page) {
	form := f.ReceiveForm(req)
	id := form.Get("id")
	dn := form.Get("displayName")
	pwd := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

	sql.AddAccount(id, dn, pwd)
	f.SendNavigate(res, "login")
}

func Register() (
	page string,
	show f.PageFunction,
	action f.PageFunction,
) {
	page = "register"
	action = registerAction
	return
}
