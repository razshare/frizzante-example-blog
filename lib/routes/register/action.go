package register

import (
	"context"
	"crypto/sha256"
	"fmt"
	"main/lib/core/clients"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/database"
	"main/lib/database/sqlc"
)

func Action(client *clients.Client) {
	var err error
	var id = receive.FormValue(client, "id")
	var displayName = receive.FormValue(client, "displayName")
	var password = receive.FormValue(client, "password")
	var hash string

	if id == "" || displayName == "" || password == "" {
		send.Navigate(client, "/register?error=please fill all fields")
		return
	}

	hash = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))

	if _, err = database.Queries.FindAccountById(context.Background(), id); err == nil {
		send.Navigatef(client, "/register?error=account %s already exists", id)
		return
	}

	if err = database.Queries.AddAccount(context.Background(), sqlc.AddAccountParams{
		ID:          id,
		DisplayName: displayName,
		Password:    hash,
	}); err != nil {
		send.Navigatef(client, "/register?error=%s", err)
		return
	}

	send.Navigate(client, "/login")
}
