package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/views"
)

func Expired(connection *connections.Connection) {
	connection.SendView(views.View{Name: "Expired"})
}
