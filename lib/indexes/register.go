package indexes

import (
	"crypto/sha256"
	"fmt"
	frz "github.com/razshare/frizzante"
	"main/lib"
	"main/lib/sql"
)

func registerAndRedirect(req *frz.Request, res *frz.Response, p *frz.Page) {
	form := frz.ReceiveForm(req)

	id := form.Get("id")
	displayName := form.Get("display_name")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

	sql.AddAccount(id, displayName, password)
	frz.SendRedirectToPage(res, "login", lib.NoProps)
}

func Register() (
	show frz.PageFunction,
	action frz.PageFunction,
) {
	action = registerAndRedirect
	return
}
