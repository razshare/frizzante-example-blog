package logout

import (
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/session"
)

func Action(c *client.Client) {
	s := session.Start(receive.SessionId(c))
	s.Verified = false
	send.Navigate(c, "/login")
}
