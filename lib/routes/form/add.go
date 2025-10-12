package form

import (
	"context"
	"main/lib/core/clients"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/database"
	"main/lib/database/sqlc"
	"main/lib/security"
	"main/lib/sessions"
)

func Add(client *clients.Client) {
	session := sessions.Start(receive.SessionId(client))

	var form sqlc.AddArticleParams
	if !receive.Form(client, &form) {
		session.FormError = "could not parse form"
		send.Navigate(client, "/form")
		return
	}

	if form.Title == "" {
		session.FormError = "article title cannot be empty"
		send.Navigate(client, "/form")
		return
	}

	if form.Content == "" {
		session.FormError = "article content cannot be empty"
		send.Navigate(client, "/form")
		return
	}

	form.ID = security.RandomHex(36)

	if err := database.Queries.AddArticle(context.Background(), form); err != nil {
		session.FormError = err.Error()
		send.Navigate(client, "/form")
		return
	}

	send.Navigate(client, "/board")
}
