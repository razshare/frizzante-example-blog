package handlers

import "github.com/razshare/frizzante/connections"

func Default(con *connections.Connection) {
	connections.SendFileOrElse(con, func() {
		connections.SendNavigate(con, "/board")
	})
}
