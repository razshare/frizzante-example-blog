package board

import (
	"github.com/razshare/frizzante/client"
	"github.com/razshare/frizzante/receive"
	"github.com/razshare/frizzante/stack"
	"strconv"
)

var PageSize int64 = 10

func Paginate(c *client.Client) int64 {
	var p int64

	if pstr := receive.Query(c, "page"); pstr != "" {
		var err error
		p, err = strconv.ParseInt(pstr, 10, 64)
		if err != nil {
			c.Scope.Container.Config.ErrorLog.Println(err, stack.Trace())
			return 0
		}
	}

	if p <= 0 {
		p = 0
	}

	return p
}
