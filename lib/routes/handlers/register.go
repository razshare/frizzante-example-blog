package handlers

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/views"
	"main/lib/database"
	"main/lib/database/sqlc"
)

func Register(connection *connections.Connection) {
	connection.SendView(views.View{Name: "Register"})
}

func RegisterAction(connection *connections.Connection) {
	form := connection.ReceiveForm()
	id := form.Get("id")

	displayName := form.Get("displayName")
	rawPassword := form.Get("password")

	if "" == id || "" == displayName || "" == rawPassword {
		connection.SendView(views.View{Name: "Register", Data: map[string]any{
			"error": "please fill all fields",
		}})
		return
	}

	password := fmt.Sprintf("%x", sha256.Sum256([]byte(rawPassword)))

	_, accountError := database.Queries.FindAccountById(context.Background(), id)
	if nil == accountError {
		connection.SendView(views.View{Name: "Register", Data: map[string]any{
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
		connection.SendView(views.View{Name: "Register", Data: map[string]any{
			"error": addError.Error(),
		}})
		return
	}

	connection.SendNavigate("/login")
}
