package handlers

import "github.com/razshare/frizzante/frz"

func Login(c *frz.Connection) {
	c.SendView(frz.View{Name: "Login"})
}
