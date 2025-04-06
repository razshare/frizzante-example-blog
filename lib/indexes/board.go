package indexes

import (
	f "github.com/razshare/frizzante"
)

func showIfVerified(req *f.Request, res *f.Response, _ *f.Page) {
	get, _, _ := f.SessionStart(req, res)
	if !get("verified", false).(bool) {
		f.SendNavigate(res, "login")
	}
}

func Board() (
	show f.PageFunction,
	action f.PageFunction,
) {
	show = showIfVerified
	return
}
