package form

import (
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/core/view"
)

func View(client *client.Client) {
	send.View(client, view.View{
		Name: "ArticleForm",
		Props: map[string]string{
			"error": receive.Query(client, "error"),
		},
	})
}
