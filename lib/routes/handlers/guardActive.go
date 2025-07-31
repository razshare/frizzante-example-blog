package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"main/lib/state"
	"time"
)

func GuardActive(connection *connections.Connection, allow func()) {
	session := sessions.New(connection, state.State{}).Start()
	defer session.Save()

	if time.Since(session.State.LastActivity) > 30*time.Minute {
		connection.SendNavigate("/expired")
		return
	}

	session.State.LastActivity = time.Now()
	allow()
}
