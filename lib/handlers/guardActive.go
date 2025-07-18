package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"main/lib"
	"time"
)

func GuardActive(con *connections.Connection, allow func()) {
	session := sessions.StartEmpty[lib.State](con)
	defer sessions.Save(session)

	if time.Since(session.State.LastActivity) > 30*time.Minute {
		connections.SendNavigate(con, "/expired")
		return
	}

	session.State.LastActivity = time.Now()
	allow()
}
