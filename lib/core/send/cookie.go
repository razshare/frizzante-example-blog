package send

import (
	"fmt"
	"net/url"

	"main/lib/core/client"
)

// Cookie sends a cookies to the client.
func Cookie(client *client.Client, key string, value string) {
	Header(client, "Set-Cookie", fmt.Sprintf("%s=%s; Path=/; HttpOnly", url.QueryEscape(key), url.QueryEscape(value)))
}
