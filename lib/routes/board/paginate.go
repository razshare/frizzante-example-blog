package board

import (
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/stack"
	"strconv"
)

var PageSize int64 = 10

func Paginate(client *client.Client) int64 {
	var page int64
	var pageQuery string
	var err error

	if pageQuery = receive.Query(client, "page"); pageQuery != "" {
		if page, err = strconv.ParseInt(pageQuery, 10, 64); err != nil {
			client.Config.ErrorLog.Println(err, stack.Trace())
			return 0
		}
	}

	if page <= 0 {
		page = 0
	}

	return page
}
