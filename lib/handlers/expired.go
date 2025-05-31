package handlers

import frz "github.com/razshare/frizzante"

func GetExpired(c *frz.Connection) {
	c.SendView(frz.View{Name: "Expired"})
}
