package guards

import (
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/session"
)

func Verified(c *client.Client, allow func()) {
	s := session.Start(receive.SessionId(c))

	if !s.Verified {
		send.Navigate(c, "/login")
		return
	}

	allow()
}
