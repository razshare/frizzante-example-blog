package is_not_expired

import (
	"main/lib/core/clients"
	"main/lib/core/guards"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/sessions"
	"time"
)

var Guard = guards.Guard{Name: "is-not-expired", Handler: func(client *clients.Client, allow func()) {
	state := sessions.Start(receive.SessionId(client))

	if time.Since(state.LastActivity) > 30*time.Minute {
		send.Navigate(client, "/expired?error=session expired")
		return
	}

	state.LastActivity = time.Now()

	allow()
}}
