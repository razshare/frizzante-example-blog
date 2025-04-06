package indexes

import (
	"crypto/sha256"
	"fmt"
	f "github.com/razshare/frizzante"
	"main/lib/sql"
)

func loginAndRedirect(req *f.Request, res *f.Response, p *f.Page) {
	_, set, _ := f.SessionStart(req, res)

	form := f.ReceiveForm(req)
	id := form.Get("id")
	pwd := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

	if !sql.VerifyAccount(id, pwd) {
		f.PageWithData(p, "error", "Invalid credentials")
		return
	}
	set("verified", true)

	f.SendNavigateWithParameters(res, "board", map[string]string{
		"user": "asd",
	})
}

func Login() (
	show f.PageFunction,
	action f.PageFunction,
) {
	action = loginAndRedirect
	return
}
