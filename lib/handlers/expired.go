package handlers

import "github.com/razshare/frizzante"

func GetExpired(c *frizzante.Connection) {
	c.SendView(frizzante.View{Name: "Expired"})
}
