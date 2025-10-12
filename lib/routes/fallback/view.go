package fallback

import (
	"main/lib/core/clients"
	"main/lib/core/send"
)

func View(client *clients.Client) {
	if !send.RequestedFile(client) {
		send.Navigate(client, "/board")
	}
}
