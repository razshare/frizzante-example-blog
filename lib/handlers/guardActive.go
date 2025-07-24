package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"main/lib"
	"time"
)

func GuardActive(con *connections.Connection, allow func()) {
	session := sessions.Start[lib.State](con)
	defer session.Save()

	if time.Since(session.State.LastActivity) > 30*time.Minute {
		con.SendNavigate("/expired")
		return
	}

	session.State.LastActivity = time.Now()
	allow()
}
