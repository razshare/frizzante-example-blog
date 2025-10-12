package board

import (
	"main/lib/core/clients"
	"main/lib/core/receive"
	"main/lib/core/stack"
	"strconv"
)

var PageSize int64 = 4

func Paginate(client *clients.Client) int64 {
	var err error
	var page int64

	if query := receive.Query(client, "page"); query != "" {
		if page, err = strconv.ParseInt(query, 10, 64); err != nil {
			client.Config.ErrorLog.Println(err, stack.Trace())
			return 0
		}
	}

	if page <= 0 {
		page = 0
	}

	return page
}
