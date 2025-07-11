package handlers

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib"
	"main/lib/database"
	"main/lib/database/sqlc"
	"time"
)

func Login(con *connections.Connection) {
	con.SendView(views.View{Name: "Login"})
}

func LoginAction(con *connections.Connection) {
	session := sessions.StartEmpty[lib.State](con)
	defer session.Save()

	form := con.ReceiveForm()
	id := form.Get("id")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

	_, accountError := database.Queries.VerifyAccount(context.Background(), sqlc.VerifyAccountParams{
		ID:       id,
		Password: password,
	})

	if nil != accountError {
		con.SendView(views.View{Name: "Login", Data: map[string]any{
			"error": "invalid credentials",
		}})
		return
	}

	session.State.LastActivity = time.Now()
	session.State.Verified = true
	session.State.AccountId = id

	con.SendNavigate("/board")
}
