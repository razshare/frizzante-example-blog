package guards

import (
	"github.com/razshare/frizzante/frz"
	"main/lib"
)

func Verified(c *frz.Connection, allow func()) {
	state, _ := frz.Session(c, lib.State{})

	if !state.Verified {
		c.SendNavigate("/login")
		return
	}

	allow()
}
