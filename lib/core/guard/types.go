package guard

import (
	"main/lib/core/client"
	"main/lib/core/tag"
)

type Guard struct {
	Name    string
	Handler func(client *client.Client, allow func())
	Tags    []tag.Tag
}
