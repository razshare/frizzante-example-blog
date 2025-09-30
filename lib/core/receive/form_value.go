package receive

import (
	"errors"
	"net/http"

	"main/lib/core/client"
	"main/lib/core/stack"
)

// FormValue reads the first form value associated with the given key and returns it.
func FormValue(client *client.Client, key string) string {
	if client.WebSocket != nil {
		client.Config.ErrorLog.Println("web socket connections cannot parse forms", stack.Trace())
		return ""
	}

	if client.Request.Form == nil {
		if err := client.Request.ParseMultipartForm(MaxFormSize); err != nil {
			if !errors.Is(err, http.ErrNotMultipart) {
				return ""
			}
		}
	}

	return client.Request.Form.Get(key)
}
