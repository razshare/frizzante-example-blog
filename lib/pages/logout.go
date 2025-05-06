package pages

import f "github.com/razshare/frizzante"

func Logout(context f.PageContext) {
	// Context.
	path, view, base, action := context()

	// Configure.
	path("/logout")
	view(f.ViewReference("Logout"))
	base(func(_ *f.Request, response *f.Response, _ *f.View) {
		f.ResponseSendNavigate(response, "Login")
	})
	action(func(request *f.Request, response *f.Response, _ *f.View) {
		_, set, _ := f.SessionStart(request, response)
		set("verified", false)
		f.ResponseSendNavigate(response, "Login")
	})
}
