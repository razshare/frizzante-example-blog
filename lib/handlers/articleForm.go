package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/views"
)

func ArticleForm(con *connections.Connection) {
	con.SendView(views.View{Name: "ArticleForm"})
}
