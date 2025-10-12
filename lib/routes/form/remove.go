package article

import (
	"main/lib/core/clients"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/database"
)

func Remove(client *clients.Client) {
	if err := database.Queries.RemoveArticle(client.Request.Context(), receive.Query(client, "id")); err != nil {
		send.Navigatef(client, "/board?error=%s", err)
		return
	}
	send.Navigate(client, "/board")
}
