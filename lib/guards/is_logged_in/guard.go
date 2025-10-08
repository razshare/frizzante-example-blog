package is_logged_in

import (
	"main/lib/core/clients"
	"main/lib/core/guards"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/sessions"
)

var Guard = guards.Guard{Name: "is-logged-in", Handler: func(client *clients.Client, allow func()) {
	state := sessions.Start(receive.SessionId(client))

	if !state.LoggedIn {
		send.Navigate(client, "/login")
		return
	}

	allow()
}}
