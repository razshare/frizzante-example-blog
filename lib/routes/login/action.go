package login

import (
	"context"
	"crypto/sha256"
	"fmt"
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/database/sqlite"
	"main/lib/database/sqlite/sqlc"
	"main/lib/session/memory"
	"time"
)

func Action(client *client.Client) {
	state := memory.Start(receive.SessionId(client))
	id := receive.FormValue(client, "id")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(receive.FormValue(client, "password"))))

	if _, err := sqlite.Queries.VerifyAccount(context.Background(), sqlc.VerifyAccountParams{
		ID:       id,
		Password: password,
	}); err != nil {
		send.Navigate(client, "/login?error=invalid credentials")
		return
	}

	state.LastActivity = time.Now()
	state.LoggedIn = true
	state.AccountId = id

	send.Navigate(client, "/board")
}
