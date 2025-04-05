package indexes

import (
	"crypto/sha256"
	"fmt"
	frz "github.com/razshare/frizzante"
	"main/lib/sql"
)

func loginAndRedirect(req *frz.Request, res *frz.Response, p *frz.Page) {
	_, set, _ := frz.SessionStart(req, res)

	form := frz.ReceiveForm(req)
	id := form.Get("id")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

	if !sql.VerifyAccount(id, password) {
		frz.PageWithData(p, "error", "Invalid credentials")
		return
	}
	set("verified", true)

	frz.SendNavigateWithParameters(res, "board", map[string]string{
		"user": "asd",
	})
}

func Login() (
	show frz.PageFunction,
	action frz.PageFunction,
) {
	action = loginAndRedirect
	return
}
