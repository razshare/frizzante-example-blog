package guards

import (
	"github.com/razshare/frizzante"
	"main/lib/sessions"
	"time"
)

func NotExpired(request *frizzante.Request, response *frizzante.Response) bool {
	session := frizzante.SessionStart(request, response, sessions.Adapter)

	if time.Since(session.Data.LastActivity) > 30*time.Minute {
		session.Destroy()
		response.SendNavigate("/expired")
		return false
	}

	session.Data.LastActivity = time.Now()
	session.Save()
	return true
}
