package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"main/lib/state"
)

func LogoutAction(connection *connections.Connection) {
	session := sessions.New(connection, state.State{}).Start()
	defer session.Save()

	session.State.Verified = false
	connection.SendNavigate("/login")
}
