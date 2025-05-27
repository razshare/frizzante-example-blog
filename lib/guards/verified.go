package guards

import (
	"github.com/razshare/frizzante"
	"main/lib/sessions"
)

func Verified(request *frizzante.Request, response *frizzante.Response) bool {
	session := frizzante.SessionStart(request, response, sessions.Adapter)
	verified := session.Data.Verified
	session.Save()

	if !verified {
		response.SendNavigate("/login")
		return false
	}

	return true
}
