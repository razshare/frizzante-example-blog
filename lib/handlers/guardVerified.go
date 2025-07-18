package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"main/lib"
)

func GuardVerified(con *connections.Connection, allow func()) {
	session := sessions.StartEmpty[lib.State](con)
	defer sessions.Save(session)

	if !session.State.Verified {
		connections.SendNavigate(con, "/login")
		return
	}
	allow()
}
