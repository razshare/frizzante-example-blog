package guards

import (
	frz "github.com/razshare/frizzante"
	"main/lib"
	"time"
)

func Active(c *frz.Connection, allow func()) {
	state, operator := frz.Session(c, lib.State{})
	defer operator.Save(state)

	if time.Since(state.LastActivity) > 30*time.Minute {
		c.SendNavigate("/expired")
		return
	}

	state.LastActivity = time.Now()
	allow()
}
