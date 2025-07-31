package handlers

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib/database"
	"main/lib/database/sqlc"
	"main/lib/state"
	"time"
)

func Login(connection *connections.Connection) {
	connection.SendView(views.View{Name: "Login"})
}

func LoginAction(connection *connections.Connection) {
	session := sessions.Start(connection, state.State{})
	defer session.Save()

	form := connection.ReceiveForm()
	id := form.Get("id")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

	_, accountError := database.Queries.VerifyAccount(context.Background(), sqlc.VerifyAccountParams{
		ID:       id,
		Password: password,
	})

	if nil != accountError {
		connection.SendView(views.View{Name: "Login", Data: map[string]any{
			"error": "invalid credentials",
		}})
		return
	}

	session.State.LastActivity = time.Now()
	session.State.Verified = true
	session.State.AccountId = id

	connection.SendNavigate("/board")
}
