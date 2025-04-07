package indexes

import f "github.com/razshare/frizzante"

func logoutShowFunction(_ *f.Request, res *f.Response, _ *f.Page) {
	f.SendNavigate(res, "login")
}

func logoutActionFunction(req *f.Request, res *f.Response, _ *f.Page) {
	_, set, _ := f.SessionStart(req, res)
	set("verified", false)
	f.SendNavigate(res, "login")
}

func Logout(
	route func(path string, page string),
	show func(showFunction f.PageFunction),
	action func(actionFunction f.PageFunction),
) {
	route("/logout", "logout")
	show(logoutShowFunction)
	action(logoutActionFunction)
}
