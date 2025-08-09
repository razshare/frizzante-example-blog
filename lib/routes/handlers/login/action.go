package login

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/razshare/frizzante/client"
	"github.com/razshare/frizzante/receive"
	"github.com/razshare/frizzante/send"
	"github.com/razshare/frizzante/view"
	"main/lib/database"
	"main/lib/database/sqlc"
	"main/lib/session"
	"time"
)

func Action(c *client.Client) {
	s := session.Start(receive.SessionId(c))

	f := receive.Form(c)

	id := f.Get("id")
	psw := fmt.Sprintf("%x", sha256.Sum256([]byte(f.Get("password"))))

	_, err := database.Queries.VerifyAccount(context.Background(), sqlc.VerifyAccountParams{
		ID:       id,
		Password: psw,
	})

	if err != nil {
		send.View(c, view.View{Name: "Login", Data: map[string]any{
			"error": "invalid credentials",
		}})
		return
	}

	s.LastActivity = time.Now()
	s.Verified = true
	s.AccountId = id

	send.Navigate(c, "/board")
}
