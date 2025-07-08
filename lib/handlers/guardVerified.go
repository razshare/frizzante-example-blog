package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"main/lib"
)

func GuardVerified(con *connections.Connection, allow func()) {
	session := sessions.StartEmpty[lib.State](con)
	defer session.Save()

	if !session.State.Verified {
		con.SendNavigate("/login")
		return
	}
	allow()
}
