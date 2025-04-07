package indexes

import f "github.com/razshare/frizzante"

func logoutAction(req *f.Request, res *f.Response, _ *f.Page) {
	_, set, _ := f.SessionStart(req, res)
	set("verified", false)
	f.SendNavigate(res, "login")
}

func Logout() (
	page string,
	show f.PageFunction,
	action f.PageFunction,
) {
	page = "logout"
	action = logoutAction
	return
}
