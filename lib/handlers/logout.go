package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"main/lib"
)

func LogoutAction(con *connections.Connection) {
	session := sessions.New(con, lib.State{}).Start()
	defer session.Save()

	session.State.Verified = false
	con.SendNavigate("/login")
}
