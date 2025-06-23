package guards

import (
	"github.com/razshare/frizzante/libcon"
	"github.com/razshare/frizzante/libsession"
	"main/lib"
)

func Verified(con *libcon.Connection, allow func()) {
	state, _ := libsession.Session(con, lib.State{})

	if !state.Verified {
		con.SendNavigate("/login")
		return
	}

	allow()
}
