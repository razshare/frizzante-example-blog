package indexes

import (
	frz "github.com/razshare/frizzante"
	"main/lib"
)

func logoutAndRedirect(req *frz.Request, res *frz.Response, p *frz.Page) {
	_, set, _ := frz.SessionStart(req, res)
	set("verified", false)
	frz.SendRedirectToPage(res, "login", lib.NoProps)
}

func Logout() (
	show frz.PageFunction,
	action frz.PageFunction,
) {
	action = logoutAndRedirect
	return
}
