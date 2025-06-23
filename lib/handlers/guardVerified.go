package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"main/lib"
)

func GuardVerified(con *connections.Connection, allow func()) {
	state, operator := sessions.StartEmpty[lib.State](con)
	defer operator.Save(state)

	if !state.Verified {
		con.SendNavigate("/login")
		return
	}
	allow()
}
