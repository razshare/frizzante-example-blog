package routes

import (
	"github.com/razshare/frizzante"
	"main/lib"
	"main/lib/sessions"
)

func init() {
	lib.Server.WithRequestHandler("GET /logout", GetLogout)
}

func GetLogout(req *frizzante.Request, res *frizzante.Response) {
	session := frizzante.SessionStart(req, res, sessions.Adapter)
	session.Data.Verified = false
	session.Save()
	res.SendNavigate("/login")
}
