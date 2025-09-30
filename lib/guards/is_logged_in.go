package guards

import (
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/session/memory"
)

func IsLoggedIn(client *client.Client, allow func()) {
	state := memory.Start(receive.SessionId(client))

	if !state.LoggedIn {
		send.Navigate(client, "/login")
		return
	}

	allow()
}
