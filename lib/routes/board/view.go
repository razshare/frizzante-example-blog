package board

import (
	"context"
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/core/view"
	"main/lib/database"
	"main/lib/database/sqlc"
	"main/lib/session/memory"
)

func View(client *client.Client) {
	var page = Paginate(client)
	var articles []sqlc.Article
	var count int
	var hasMore bool
	var err error

	if articles, err = database.Queries.FindArticles(
		context.Background(),
		sqlc.FindArticlesParams{
			Offset: PageSize * page,

			// Get an additional element.
			// If this element is present in the result,
			// then it means the next page is available.
			// We will then shave this element away when
			// sending the slice to the view.
			Limit: PageSize + 1,
		},
	); err != nil {
		send.Navigatef(client, "/board?error=%s", err.Error())
		return
	}

	count = len(articles)
	hasMore = count == int(PageSize)+1
	state := memory.Start(receive.SessionId(client))

	if articles = articles[:count]; articles == nil {
		articles = make([]sqlc.Article, 0)
	}

	// Send the views.
	send.View(client, view.View{Name: "Board", Props: map[string]any{
		"verified": state.Verified,
		"expired":  state.Expired,
		"page":     page,
		"hasMore":  hasMore,
		"articles": articles,
	}})
}
