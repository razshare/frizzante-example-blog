package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"main/lib"
)

func GuardVerified(con *connections.Connection, allow func()) {
	state, _ := sessions.Start[lib.State](con)
	if !state.Verified {
		con.SendNavigate("/login")
		return
	}
	allow()
}
