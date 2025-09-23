package article

import (
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/database"
)

func Remove(client *client.Client) {
	if err := database.Queries.RemoveArticle(client.Request.Context(), receive.Query(client, "id")); err != nil {
		send.Navigatef(client, "/board?error=%s", err.Error())
		return
	}
	send.Navigate(client, "/board")
}
