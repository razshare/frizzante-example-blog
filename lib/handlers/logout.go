package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"main/lib"
)

func LogoutAction(con *connections.Connection) {
	session := sessions.StartEmpty[lib.State](con)
	defer sessions.Save(session)

	session.State.Verified = false
	connections.SendNavigate(con, "/login")
}
