package register

import (
	"main/lib/core/clients"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/core/types"
	"main/lib/core/views"
	"main/lib/sessions"
)

func init() {
	_ = types.Generate[Props]()
}

type Props struct {
	Error string `json:"error"`
}

func View(client *clients.Client) {
	session := sessions.Start(receive.SessionId(client))
	defer func() { session.RegisterError = "" }()
	send.View(client, views.View{Name: "Register", Props: Props{
		Error: session.RegisterError,
	}})
}
