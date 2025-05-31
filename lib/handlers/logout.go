package handlers

import (
	"github.com/razshare/frizzante"
	"main/lib"
)

func GetLogout(c *frizzante.Connection) {
	lib.SessionStartPublic(c, func(state *lib.State) {
		state.Verified = false
		c.SendNavigate("/login")
	})
}
