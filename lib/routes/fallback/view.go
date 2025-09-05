package fallback

import (
	"main/lib/core/client"
	"main/lib/core/send"
)

func View(c *client.Client) {
	send.FileOrElse(c, func() {
		send.Navigate(c, "/board")
	})
}
