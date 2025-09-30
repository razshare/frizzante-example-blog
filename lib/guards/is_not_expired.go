package guards

import (
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/session/memory"
	"time"
)

func IsNotExpired(client *client.Client, allow func()) {
	state := memory.Start(receive.SessionId(client))

	if time.Since(state.LastActivity) > 30*time.Minute {
		send.Navigate(client, "/expired?error=session expired")
		return
	}

	state.LastActivity = time.Now()

	allow()
}
