package handlers

import "github.com/razshare/frizzante/frz"

func Register(c *frz.Connection) {
	c.SendView(frz.View{Name: "Register"})
}
