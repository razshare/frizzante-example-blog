package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"main/lib"
)

func GuardVerified(con *connections.Connection, allow func()) {
	session := sessions.New(con, lib.State{}).Start()
	defer session.Save()

	if !session.State.Verified {
		con.SendNavigate("/login")
		return
	}
	allow()
}
