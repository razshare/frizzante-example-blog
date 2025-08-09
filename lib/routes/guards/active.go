package guards

import (
	"github.com/razshare/frizzante/client"
	"github.com/razshare/frizzante/receive"
	"github.com/razshare/frizzante/send"
	"main/lib/session"
	"time"
)

func Active(c *client.Client, allow func()) {
	s := session.Start(receive.SessionId(c))

	if time.Since(s.LastActivity) > 30*time.Minute {
		send.Navigate(c, "/expired")
		return
	}

	s.LastActivity = time.Now()

	allow()
}
