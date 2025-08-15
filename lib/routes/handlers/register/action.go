package register

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
)

func Action(c *client.Client) {
	f := receive.Form(c)

	id := f.Get("id")
	n := f.Get("displayName")
	psw := f.Get("password")

	if id == "" || n == "" || psw == "" {
		send.View(c, view.View{Name: "Register", Props: map[string]any{
			"error": "please fill all fields",
		}})
		return
	}

	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(psw)))

	_, err := database.Queries.FindAccountById(context.Background(), id)
	if err != nil {
		send.View(c, view.View{Name: "Register", Props: map[string]any{
			"error": fmt.Sprintf("account `%s` already exists", id),
		}})
		return
	}

	err = database.Queries.AddAccount(context.Background(), sqlc.AddAccountParams{
		ID:          id,
		DisplayName: n,
		Password:    hash,
	})

	if err != nil {
		send.View(c, view.View{Name: "Register", Props: map[string]any{
			"error": err.Error(),
		}})
		return
	}

	send.Navigate(c, "/login")
}
