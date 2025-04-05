package indexes

import (
	frz "github.com/razshare/frizzante"
	"main/lib"
)

func showIfVerified(req *frz.Request, res *frz.Response, p *frz.Page) {
	get, _, _ := frz.SessionStart(req, res)
	if !get("verified", false).(bool) {
		frz.SendRedirectToPage(res, "login", lib.NoProps)
	}
}

func Board() (
	show frz.PageFunction,
	action frz.PageFunction,
) {
	show = showIfVerified
	return
}
