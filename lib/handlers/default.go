package handlers

import "github.com/razshare/frizzante/connections"

func Default(con *connections.Connection) {
	con.SendFileOrElse(func() {
		con.SendNavigate("/board")
	})
}
