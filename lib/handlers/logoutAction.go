package handlers

import (
	"github.com/razshare/frizzante/libcon"
	"github.com/razshare/frizzante/libsession"
	"main/lib"
)

func LogoutAction(con *libcon.Connection) {
	state, operator := libsession.Session(con, lib.State{})
	defer operator.Save(state)
	state.Verified = false
	con.SendNavigate("/login")
}
