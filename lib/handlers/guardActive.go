package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"main/lib"
	"time"
)

func GuardActive(con *connections.Connection, allow func()) {
	state, operator := sessions.StartEmpty[lib.State](con)
	defer operator.Save(state)

	if time.Since(state.LastActivity) > 30*time.Minute {
		con.SendNavigate("/expired")
		return
	}

	state.LastActivity = time.Now()
	allow()
}
