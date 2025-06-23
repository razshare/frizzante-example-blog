package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"main/lib"
)

func LogoutAction(con *connections.Connection) {
	state, operator := sessions.StartEmpty[lib.State](con)
	defer operator.Save(state)

	state.Verified = false
	con.SendNavigate("/login")
}
