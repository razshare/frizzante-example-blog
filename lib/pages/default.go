package pages

import f "github.com/razshare/frizzante"

func Default(context f.PageContext) {
	// Context.
	path, _, base, action := context()

	// Configure.
	path("/")
	base(func(_ *f.Request, response *f.Response, _ *f.View) {
		f.ResponseSendNavigate(response, "Login")
	})
	action(func(_ *f.Request, response *f.Response, _ *f.View) {
		f.ResponseSendNavigate(response, "Login")
	})
}
