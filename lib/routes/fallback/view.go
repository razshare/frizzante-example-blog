package fallback

import (
	"main/lib/core/client"
	"main/lib/core/send"
	"os"
)

func View(client *client.Client) {
	send.FileOrElse(client, send.FileOrElseConfig{
		UseDisk: os.Getenv("DEV") == "1",
		OrElse:  func() { send.Navigate(client, "/board") },
	})
}
