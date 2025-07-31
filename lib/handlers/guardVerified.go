package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"main/lib/state"
)

func GuardVerified(connection *connections.Connection, allow func()) {
	session := sessions.Start(connection, state.State{})
	defer session.Save()

	if !session.State.Verified {
		connection.SendNavigate("/login")
		return
	}
	allow()
}
