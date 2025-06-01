package guards

import (
	frz "github.com/razshare/frizzante"
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
