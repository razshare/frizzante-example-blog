package fallback

import (
	"main/lib/core/clients"
	"main/lib/core/send"
)

func View(client *clients.Client) {
	send.FileOrElse(client, func() { send.Navigate(client, "/board") })
}
