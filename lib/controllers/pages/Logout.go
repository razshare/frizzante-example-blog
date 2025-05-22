package pages

import (
	f "github.com/razshare/frizzante"
	"main/lib/sessions"
)

var Logout = f.NewPageController().WithBase(logoutBase).WithAction(logoutAction)

func logoutBase(req *f.Request, res *f.Response) {
	session := f.SessionStart(req, res, sessions.Archived)
	session.Data.Verified = false
	session.Save()
	res.SendNavigate("Login")
}

func logoutAction(req *f.Request, res *f.Response) {
	session := f.SessionStart(req, res, sessions.Archived)
	session.Data.Verified = false
	session.Save()
	res.SendNavigate("Login")
}
