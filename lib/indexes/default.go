package indexes

import f "github.com/razshare/frizzante"

func defaultShow(_ *f.Request, res *f.Response, _ *f.Page) {
	// Show page.
	f.SendNavigate(res, "login")
}

func defaultAction(_ *f.Request, res *f.Response, _ *f.Page) {
	// Run page action.
	f.SendNavigate(res, "login")
}

func Default() (
	page string,
	show f.PageFunction,
	action f.PageFunction,
) {
	page = "default /"
	show = defaultShow
	action = defaultAction
	return
}
