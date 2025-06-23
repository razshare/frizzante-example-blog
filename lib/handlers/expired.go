package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/views"
)

func Expired(con *connections.Connection) {
	con.SendView(views.View{Name: "Expired"})
}
