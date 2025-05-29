package routes

import (
	"github.com/razshare/frizzante"
	"main/lib/sessions"
)

func GetLogout(req *frizzante.Request, res *frizzante.Response) {
	session := frizzante.SessionStart(req, res, sessions.Adapter)
	session.Data.Verified = false
	session.Save()
	res.SendNavigate("/login")
}
