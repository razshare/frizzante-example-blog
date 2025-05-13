package guards

import (
	f "github.com/razshare/frizzante"
	"time"
)

func Verified(request *f.Request, response *f.Response, pass func()) {
	session := f.SessionStart(request, response)

	verified := f.SessionHas(session, "verified") && f.SessionGetBool(session, "verified")

	if !verified {
		f.ResponseSendNavigate(response, "Login")
		return
	}

	if !f.SessionHas(session, "lastActivity") {
		f.SessionSetTime(session, "lastActivity", time.Now())
	}

	lastActivity := f.SessionGetTime(session, "lastActivity")

	if time.Since(lastActivity) > 30*time.Minute {
		f.SessionDestroy(session)
		f.ResponseSendNavigate(response, "Expired")
		return
	}

	f.SessionSetTime(session, "lastActivity", time.Now())

	pass()
}
