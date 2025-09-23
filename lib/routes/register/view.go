package register

import (
	"main/lib/core/client"
	"main/lib/core/send"
	"main/lib/core/view"
)

func View(client *client.Client) {
	send.View(client, view.View{Name: "Register"})
}
