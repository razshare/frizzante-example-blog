package expired

import (
	"main/lib/core/client"
	"main/lib/core/send"
	"main/lib/core/view"
)

func View(c *client.Client) {
	send.View(c, view.View{Name: "Expired", Props: map[string]any{
		"error": "",
	}})
}
