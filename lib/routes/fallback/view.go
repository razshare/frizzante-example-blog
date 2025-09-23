package fallback

import (
	"main/lib/core/client"
	"main/lib/core/send"
)

func View(client *client.Client) {
	send.FileOrElse(client, func() { send.Navigate(client, "/board") })
}
