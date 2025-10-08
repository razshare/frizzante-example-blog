package logout

import (
	"main/lib/core/clients"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/sessions"
)

func Action(client *clients.Client) {
	session := sessions.Start(receive.SessionId(client))
	session.LoggedIn = false
	send.Navigate(client, "/login")
}
