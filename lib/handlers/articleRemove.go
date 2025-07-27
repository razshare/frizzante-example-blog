package handlers

import (
	"context"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/views"
	"main/lib/database"
)

func ArticleRemove(con *connections.Connection) {
	id := con.ReceiveQuery("id")
	removeError := database.Queries.RemoveArticle(context.Background(), id)
	if removeError != nil {
		con.SendView(views.View{Name: "Board", Data: map[string]any{
			"error": removeError,
		}})
		return
	}

	con.SendNavigate("/board")
}
