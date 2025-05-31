package handlers

import (
	frz "github.com/razshare/frizzante"
	"main/lib"
)

func GetLogout(c *frz.Connection) {
	state, operator := frz.Session(c, lib.State{})
	defer operator.Save(state)
	state.Verified = false
	c.SendNavigate("/login")
}
