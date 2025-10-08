package login

import (
	"context"
	"crypto/sha256"
	"fmt"
	"main/lib/core/clients"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/database"
	"main/lib/database/sqlc"
	"main/lib/sessions"
	"time"
)

func Action(client *clients.Client) {
	session := sessions.Start(receive.SessionId(client))
	id := receive.FormValue(client, "id")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(receive.FormValue(client, "password"))))

	if _, err := database.Queries.VerifyAccount(context.Background(), sqlc.VerifyAccountParams{
		ID:       id,
		Password: password,
	}); err != nil {
		send.Navigate(client, "/login?error=invalid credentials")
		return
	}

	session.LastActivity = time.Now()
	session.LoggedIn = true
	session.AccountId = id

	send.Navigate(client, "/board")
}
