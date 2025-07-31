package handlers

import (
	"context"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/views"
	"main/lib/database"
)

func ArticleRemove(connection *connections.Connection) {
	id := connection.ReceiveQuery("id")
	removeError := database.Queries.RemoveArticle(context.Background(), id)
	if removeError != nil {
		connection.SendView(views.View{Name: "Board", Data: map[string]any{
			"error": removeError,
		}})
		return
	}

	connection.SendNavigate("/board")
}
