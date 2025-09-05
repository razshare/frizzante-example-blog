package board

import (
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/stack"
	"strconv"
)

var PageSize int64 = 10

func Paginate(c *client.Client) int64 {
	var p int64

	if pstr := receive.Query(c, "page"); pstr != "" {
		var err error
		p, err = strconv.ParseInt(pstr, 10, 64)
		if err != nil {
			c.Config.ErrorLog.Println(err, stack.Trace())
			return 0
		}
	}

	if p <= 0 {
		p = 0
	}

	return p
}
