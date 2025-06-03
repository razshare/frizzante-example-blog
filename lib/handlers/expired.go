package handlers

import "github.com/razshare/frizzante/frz"

func GetExpired(c *frz.Connection) {
	c.SendView(frz.View{Name: "Expired"})
}
