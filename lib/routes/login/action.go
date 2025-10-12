package login

import (
	"context"
	"main/lib/core/clients"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/database"
	"main/lib/database/sqlc"
	"main/lib/security"
	"main/lib/sessions"
	"time"
)

func Action(client *clients.Client) {
	session := sessions.Start(receive.SessionId(client))

	var form sqlc.VerifyAccountParams
	if !receive.Form(client, &form) {
		session.LoginError = "could not parse form"
		send.Navigate(client, "/login")
		return
	}

	form.Password = security.Sha256(form.Password)

	if _, err := database.Queries.VerifyAccount(context.Background(), form); err != nil {
		session.LoginError = "invalid credentials"
		send.Navigate(client, "/login")
		return
	}

	session.LastActivity = time.Now()
	session.LoggedIn = true
	session.AccountId = form.ID

	send.Navigate(client, "/board")
}
