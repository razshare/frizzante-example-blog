package logout

import (
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/session/memory"
)

func Action(client *client.Client) {
	state := memory.Start(receive.SessionId(client))
	state.LoggedIn = false
	send.Navigate(client, "/login")
}
