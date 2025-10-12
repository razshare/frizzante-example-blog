package register

import (
	"context"
	"fmt"
	"main/lib/core/clients"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/database"
	"main/lib/database/sqlc"
	"main/lib/security"
	"main/lib/sessions"
)

func Action(client *clients.Client) {
	session := sessions.Start(receive.SessionId(client))

	var form sqlc.AddAccountParams
	if !receive.Form(client, &form) {
		session.RegisterError = "could not parse form"
		send.Navigate(client, "/register")
		return
	}

	if form.ID == "" || form.DisplayName == "" || form.Password == "" {
		session.RegisterError = "please fill all fields"
		send.Navigate(client, "/register")
		return
	}

	form.Password = security.Sha256(form.Password)

	if _, err := database.Queries.FindAccountById(context.Background(), form.ID); err == nil {
		session.RegisterError = fmt.Sprintf("account %s already exists", form.ID)
		send.Navigate(client, "/register")
		return
	}

	if err := database.Queries.AddAccount(context.Background(), form); err != nil {
		session.RegisterError = err.Error()
		send.Navigate(client, "/register")
		return
	}

	send.Navigate(client, "/login")
}
