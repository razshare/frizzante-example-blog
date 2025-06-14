package handlers

import "github.com/razshare/frizzante/frz"

func Expired(c *frz.Connection) {
	c.SendView(frz.View{Name: "Expired"})
}
