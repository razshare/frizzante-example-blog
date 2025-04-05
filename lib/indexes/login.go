package indexes

import (
	"crypto/sha256"
	"fmt"
	frz "github.com/razshare/frizzante"
	"main/lib"
	"main/lib/sql"
)

func loginAndRedirect(req *frz.Request, res *frz.Response, p *frz.Page) {
	form := frz.ReceiveForm(req)
	_, set, _ := frz.SessionStart(req, res)

	id := form.Get("id")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

	if !sql.VerifyAccount(id, password) {
		frz.PageWithData(p, "error", "Invalid credentials")
		return
	}
	set("verified", true)
	frz.SendRedirectToPage(res, "board", lib.NoProps)
}

func Login() (
	show frz.PageFunction,
	action frz.PageFunction,
) {
	action = loginAndRedirect
	return
}
