package send

import (
	"encoding/json"

	"main/lib/core/client"
	"main/lib/core/stack"
)

// Json sends json content.
//
// Compatible with web sockets and server sent events.
func Json(client *client.Client, value any) {
	data, err := json.Marshal(value)
	if err != nil {
		client.Config.ErrorLog.Println(err, stack.Trace())
		return
	}

	if client.WebSocket == nil {
		if client.Writer.Header().Get("Content-Type") == "" {
			client.Writer.Header().Set("Content-Type", "application/json")
		}
	}

	Content(client, data)
}
