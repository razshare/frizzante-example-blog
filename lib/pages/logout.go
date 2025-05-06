package pages

import f "github.com/razshare/frizzante"

func logoutShowFunction(_ *f.Request, response *f.Response, _ *f.View) {
	f.SendNavigate(response, "login")
}

func logoutActionFunction(request *f.Request, response *f.Response, _ *f.View) {
	_, set, _ := f.SessionStart(request, response)
	set("verified", false)
	f.SendNavigate(response, "Login")
}

func Logout(context f.PageContext) {
	path, view, base, action := context()
	path("/logout")
	view(f.ViewReference("Logout"))
	base(logoutShowFunction)
	action(logoutActionFunction)
}
