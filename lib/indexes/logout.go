package indexes

import f "github.com/razshare/frizzante"

func logoutAction(req *f.Request, res *f.Response, _ *f.Page) {
	_, set, _ := f.SessionStart(req, res)
	set("verified", false)
	f.SendNavigate(res, "login")
}

func Logout() (
	s f.PageFunction,
	a f.PageFunction,
) {
	a = logoutAction
	return
}
