package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"main/lib"
)

func LogoutAction(con *connections.Connection) {
	session := sessions.Start[lib.State](con)
	defer session.Save()

	session.State.Verified = false
	con.SendNavigate("/login")
}
