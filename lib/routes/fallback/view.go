package fallback

import (
	"github.com/razshare/frizzante/client"
	"github.com/razshare/frizzante/send"
)

func View(c *client.Client) {
	send.FileOrElse(c, func() {
		send.Navigate(c, "/board")
	})
}
