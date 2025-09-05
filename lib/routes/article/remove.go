package article

import (
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/core/view"
	"main/lib/database"
)

func Remove(c *client.Client) {
	id := receive.Query(c, "id")

	err := database.Queries.RemoveArticle(c.Request.Context(), id)

	if err != nil {
		send.View(c, view.View{Name: "Board", Props: map[string]any{
			"error": err,
		}})
		return
	}

	send.Navigate(c, "/board")
}
