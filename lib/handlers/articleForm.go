package handlers

import "github.com/razshare/frizzante/frz"

func ArticleForm(c *frz.Connection) {
	c.SendView(frz.View{Name: "ArticleForm"})
}
