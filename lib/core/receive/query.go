package receive

import "main/lib/core/client"

// Query reads a query field and returns the value.
func Query(client *client.Client, key string) string {
	return client.Request.URL.Query().Get(key)
}
