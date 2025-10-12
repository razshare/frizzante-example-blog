package form

import (
	"main/lib/core/clients"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/database"
	"main/lib/sessions"
)

func Remove(client *clients.Client) {
	session := sessions.Start(receive.SessionId(client))
	if err := database.Queries.RemoveArticle(client.Request.Context(), receive.Query(client, "id")); err != nil {
		session.BoardError = err.Error()
	}
	send.Navigate(client, "/board")
}
