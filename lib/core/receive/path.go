package receive

import "main/lib/core/clients"

// Path reads a parameters fields and returns the value.
func Path(client *clients.Client, key string) string {
	return client.Request.PathValue(key)
}
