package guards

import (
	f "github.com/razshare/frizzante"
	"main/lib/sessions"
)

func SessionIsVerified(request *f.Request, response *f.Response) bool {
	session := f.SessionStart(request, response, sessions.Archived)
	verified := session.Data.Verified

	if !verified {
		response.SendNavigate("Login")
		return false
	}

	return true
}
