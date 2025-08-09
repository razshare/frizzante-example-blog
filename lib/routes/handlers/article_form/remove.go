package article_form

import (
	"context"
	"github.com/razshare/frizzante/client"
	"github.com/razshare/frizzante/receive"
	"github.com/razshare/frizzante/send"
	"github.com/razshare/frizzante/view"
	"main/lib/database"
)

func Remove(c *client.Client) {
	id := receive.Query(c, "id")

	err := database.Queries.RemoveArticle(context.Background(), id)

	if err != nil {
		send.View(c, view.View{Name: "Board", Data: map[string]any{
			"error": err,
		}})
		return
	}

	send.Navigate(c, "/board")
}
