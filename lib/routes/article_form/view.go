package article_form

import (
	"main/lib/core/client"
	"main/lib/core/send"
	"main/lib/core/view"
)

func View(c *client.Client) {
	send.View(c, view.View{Name: "ArticleForm", Props: map[string]any{
		"error": "",
	}})
}
