package guards

import (
	"github.com/razshare/frizzante/client"
	"github.com/razshare/frizzante/receive"
	"github.com/razshare/frizzante/send"
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
