package handlers

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/razshare/frizzante/libcon"
	"github.com/razshare/frizzante/libview"
	"main/lib/database"
	"main/lib/utilities/sqlc"
)

func RegisterAction(con *libcon.Connection) {
	form := con.ReceiveForm()
	id := form.Get("id")

	displayName := form.Get("displayName")
	rawPassword := form.Get("password")

	if "" == id || "" == displayName || "" == rawPassword {
		con.SendView(libview.View{Name: "Register", Data: map[string]any{
			"error": "please fill all fields",
		}})
		return
	}

	password := fmt.Sprintf("%x", sha256.Sum256([]byte(rawPassword)))

	_, accountError := database.Queries.FindAccountById(context.Background(), id)
	if nil == accountError {
		con.SendView(libview.View{Name: "Register", Data: map[string]any{
			"error": fmt.Sprintf("account `%s` already exists", id),
		}})
		return
	}

	addError := database.Queries.AddAccount(context.Background(), sqlc.AddAccountParams{
		ID:          id,
		DisplayName: displayName,
		Password:    password,
	})

	if nil != addError {
		con.SendView(libview.View{Name: "Register", Data: map[string]any{
			"error": addError.Error(),
		}})
		return
	}

	con.SendNavigate("/login")
}
