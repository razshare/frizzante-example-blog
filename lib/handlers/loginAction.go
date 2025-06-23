package handlers

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/razshare/frizzante/libcon"
	"github.com/razshare/frizzante/libsession"
	"github.com/razshare/frizzante/libview"
	"main/lib"
	"main/lib/database"
	"main/lib/utilities/sqlc"
	"time"
)

func LoginAction(con *libcon.Connection) {
	state, operator := libsession.Session(con, lib.State{})
	defer operator.Save(state)
	form := con.ReceiveForm()
	id := form.Get("id")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

	_, accountError := database.Queries.VerifyAccount(context.Background(), sqlc.VerifyAccountParams{
		ID:       id,
		Password: password,
	})

	if nil != accountError {
		con.SendView(libview.View{Name: "Login", Data: map[string]any{
			"error": "invalid credentials",
		}})
		return
	}

	state.LastActivity = time.Now()
	state.Verified = true
	state.AccountId = id

	con.SendNavigate("/board")
}
