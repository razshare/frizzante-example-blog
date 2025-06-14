package handlers

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/razshare/frizzante/frz"
	"main/lib"
	"main/lib/database"
	"main/lib/generated"
	"time"
)

func LoginAction(c *frz.Connection) {
	state, operator := frz.Session(c, lib.State{})
	defer operator.Save(state)
	form := c.ReceiveForm()
	id := form.Get("id")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

	_, accountError := database.Queries.VerifyAccount(context.Background(), generated.VerifyAccountParams{
		ID:       id,
		Password: password,
	})

	if nil != accountError {
		c.SendView(frz.View{Name: "Login", Data: map[string]any{
			"error": "invalid credentials",
		}})
		return
	}

	state.LastActivity = time.Now()
	state.Verified = true
	state.AccountId = id

	c.SendNavigate("/board")
}
