package handlers

import (
	"github.com/razshare/frizzante/frz"
	"main/lib"
)

func LogoutAction(c *frz.Connection) {
	state, operator := frz.Session(c, lib.State{})
	defer operator.Save(state)
	state.Verified = false
	c.SendNavigate("/login")
}
