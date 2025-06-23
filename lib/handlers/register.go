package handlers

import (
	"github.com/razshare/frizzante/libcon"
	"github.com/razshare/frizzante/libview"
)

func Register(con *libcon.Connection) {
	con.SendView(libview.View{Name: "Register"})
}
