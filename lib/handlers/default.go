package handlers

import "github.com/razshare/frizzante/frz"

func Default(c *frz.Connection) {
	c.SendFileOrElse(func() { Login(c) })
}
