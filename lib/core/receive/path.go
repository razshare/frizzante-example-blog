package receive

import "main/lib/core/client"

// Path reads a parameters fields and returns the value.
func Path(client *client.Client, key string) string {
	return client.Request.PathValue(key)
}
