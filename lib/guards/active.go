package guards

import (
	"github.com/razshare/frizzante/libcon"
	"github.com/razshare/frizzante/libsession"
	"main/lib"
	"time"
)

func Active(con *libcon.Connection, allow func()) {
	state, operator := libsession.Session(con, lib.State{})
	defer operator.Save(state)

	if time.Since(state.LastActivity) > 30*time.Minute {
		con.SendNavigate("/expired")
		return
	}

	state.LastActivity = time.Now()
	allow()
}
