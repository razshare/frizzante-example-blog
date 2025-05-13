package guards

import (
	f "github.com/razshare/frizzante"
	"main/lib/sessions"
	"time"
)

func Verified(request *f.Request, response *f.Response, pass func()) {
	session := f.SessionStart(request, response, sessions.Builder)

	verified := f.SessionHas(session, "verified") && f.SessionGetBool(session, "verified")

	if !verified {
		f.ResponseSendNavigate(response, "Login")
		return
	}

	if !f.SessionHas(session, "lastActivity") {
		f.SessionSetTime(session, "lastActivity", time.Now())
	}

	lastActivity := f.SessionGetTime(session, "lastActivity")

	if time.Since(lastActivity) > 20*time.Second {
		f.SessionDestroy(session)
		f.ResponseSendNavigate(response, "Expired")
		return
	}

	f.SessionSetTime(session, "lastActivity", time.Now())

	pass()
}
