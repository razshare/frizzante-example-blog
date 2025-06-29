package handlers

import (
	"context"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib"
	"main/lib/database"
	"main/lib/database/sqlc"
	"strconv"
)

var pageSize int64 = 10

func ReceivePage(con *connections.Connection) int64 {
	var page int64

	// Find page.
	if stringified := con.ReceiveQuery("page"); stringified != "" {
		var parseError error
		page, parseError = strconv.ParseInt(stringified, 10, 64)
		if parseError != nil {
			lib.Notifier.SendError(parseError)
			return 0
		}
	}

	// Guard against unusual pages.
	if page <= 0 {
		page = 0
	}
	return page
}

func Board(con *connections.Connection) {
	// Find page.
	page := ReceivePage(con)

	// Find articles.
	articles, articleError := database.Queries.FindArticles(
		context.Background(),
		sqlc.FindArticlesParams{
			Offset: pageSize * page,
			Limit:  pageSize,
		},
	)

	// Make sure "articles" is not nil.
	if articles == nil {
		articles = []sqlc.FindArticlesRow{}
	}

	// Check for errors.
	if nil != articleError {
		con.SendView(views.View{Name: "Board", Data: map[string]any{
			"error": articleError.Error(),
		}})
		return
	}

	// Find the first item in the next page.
	nextArticles, nextArticlesError := database.Queries.FindArticles(
		context.Background(),
		sqlc.FindArticlesParams{
			Offset: pageSize * (page + 1),
			Limit:  1,
		},
	)

	// Check if next page has items.
	hasMore := nextArticlesError == nil && nextArticles != nil && len(nextArticles) > 0

	state, operator := sessions.StartEmpty[lib.State](con)
	defer operator.Save(state)

	// Send the views.
	con.SendView(views.View{Name: "Board", Data: map[string]any{
		"verified": state.Verified,
		"expired":  state.Expired,
		"page":     page,
		"hasMore":  hasMore,
		"articles": articles,
	}})
}
