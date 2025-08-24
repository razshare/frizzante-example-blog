package board

import (
	"context"
	"github.com/razshare/frizzante/client"
	"github.com/razshare/frizzante/receive"
	"github.com/razshare/frizzante/send"
	"github.com/razshare/frizzante/view"
	"main/lib/database"
	"main/lib/database/sqlc"
	"main/lib/session"
)

func View(c *client.Client) {
	p := Paginate(c)

	arts, err := database.Queries.FindArticles(
		context.Background(),
		sqlc.FindArticlesParams{
			Offset: PageSize * p,

			// Get an additional element.
			// If this element is present in the result,
			// then it means the next page is available.
			// We will then shave this element away when
			// sending the slice to the view.
			Limit: PageSize + 1,
		},
	)

	if err != nil {
		send.View(c, view.View{Name: "Board", Props: map[string]any{
			"error": err.Error(),
		}})
		return
	}

	l := len(arts)

	hm := l == int(PageSize)+1

	s := session.Start(receive.SessionId(c))

	arts = arts[:l]

	if arts == nil {
		arts = make([]sqlc.Article, 0)
	}

	// Send the views.
	send.View(c, view.View{Name: "Board", Props: map[string]any{
		"verified": s.Verified,
		"expired":  s.Expired,
		"page":     p,
		"hasMore":  hm,
		"articles": arts,
		"error":    "",
	}})
}
