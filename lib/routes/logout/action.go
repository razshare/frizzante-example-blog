package logout

import (
	"github.com/razshare/frizzante/client"
	"github.com/razshare/frizzante/receive"
	"github.com/razshare/frizzante/send"
	"main/lib/session"
)

func Action(c *client.Client) {
	s := session.Start(receive.SessionId(c))
	s.Verified = false
	send.Navigate(c, "/login")
}
