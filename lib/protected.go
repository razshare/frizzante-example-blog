package lib

import (
	"github.com/razshare/frizzante"
	"time"
)

var protected = []frizzante.SessionGuard[State]{
	func(connection *frizzante.Connection, state *State, pass func()) {
		if !state.Verified {
			connection.SendNavigate("/login")
			return
		}
		pass()
	},
	func(connection *frizzante.Connection, state *State, pass func()) {
		if time.Since(state.LastActivity) > 30*time.Minute {
			connection.SendNavigate("/expired")
			return
		}

		state.LastActivity = time.Now()
		pass()
	},
}
