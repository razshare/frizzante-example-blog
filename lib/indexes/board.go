package indexes

import (
	frz "github.com/razshare/frizzante"
)

func showIfVerified(req *frz.Request, res *frz.Response, _ *frz.Page) {
	get, _, _ := frz.SessionStart(req, res)
	verified := get("verified", false).(bool)
	if !verified {
		frz.SendNavigate(res, "login")
	}
}

func Board() (
	show frz.PageFunction,
	action frz.PageFunction,
) {
	show = showIfVerified
	return
}
