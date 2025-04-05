package indexes

import frz "github.com/razshare/frizzante"

func logoutAndRedirect(req *frz.Request, res *frz.Response, _ *frz.Page) {
	_, set, _ := frz.SessionStart(req, res)
	set("verified", false)
	frz.SendNavigate(res, "login")
}

func Logout() (
	show frz.PageFunction,
	action frz.PageFunction,
) {
	action = logoutAndRedirect
	return
}
