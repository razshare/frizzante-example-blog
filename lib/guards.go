package lib

import (
	f "github.com/razshare/frizzante"
	"time"
)

func NotExpired(request *f.Request, response *f.Response) bool {
	session := f.SessionStart(request, response, SessionAdapter)

	if time.Since(session.Data.LastActivity) > 30*time.Minute {
		session.Destroy()
		response.SendNavigate("expired")
		return false
	}

	session.Data.LastActivity = time.Now()
	session.Save()
	return true
}

func Verified(request *f.Request, response *f.Response) bool {
	session := f.SessionStart(request, response, SessionAdapter)
	verified := session.Data.Verified
	session.Save()

	if !verified {
		response.SendNavigate("login")
		return false
	}

	return true
}
