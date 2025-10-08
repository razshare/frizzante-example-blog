package form

import (
	"main/lib/core/clients"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/core/types"
	"main/lib/core/views"
)

func View(client *clients.Client) {
	send.View(client, views.View{
		Name: "Form",
		Props: Props{
			Error: receive.Query(client, "error"),
		},
	})
}

type Props struct {
	Error string `json:"error"`
}

func init() {
	_ = types.Generate[Props]()
}
