package pages

import (
	f "github.com/razshare/frizzante"
	"main/lib/sessions"
)

type LogoutController struct {
	f.PageController
}

func (_ LogoutController) Configure() f.PageConfiguration {
	return f.PageConfiguration{
		Path: "/logout",
	}
}

func (_ LogoutController) Base(request *f.Request, response *f.Response) {
	session := f.SessionStart(request, response, sessions.Archived)
	session.Data.Verified = false
	session.Save()
	response.SendNavigate("Login")
}

func (_ LogoutController) Action(request *f.Request, response *f.Response) {
	session := f.SessionStart(request, response, sessions.Archived)
	session.Data.Verified = false
	session.Save()
	response.SendNavigate("Login")
}
