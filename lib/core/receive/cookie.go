package receive

import (
	"net/url"

	"main/lib/core/clients"
	"main/lib/core/stacks"
)

// Cookie reads the contents of a cookie from the message and returns the value.
func Cookie(client *clients.Client, key string) string {
	cookie, err := client.Request.Cookie(key)
	if err != nil {
		client.Config.ErrorLog.Println(err, stacks.Trace())
		return ""
	}

	data, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		client.Config.ErrorLog.Println(err, stacks.Trace())
		return ""
	}

	return data
}
