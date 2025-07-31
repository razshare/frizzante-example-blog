package handlers

import (
	"context"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/traces"
	"github.com/razshare/frizzante/views"
	"main/lib/database"
	"main/lib/database/sqlc"
	"main/lib/state"
	"strconv"
)

var pageSize int64 = 10

func ReceivePage(connection *connections.Connection) int64 {
	var page int64

	// Find page.
	if stringified := connection.ReceiveQuery("page"); stringified != "" {
		var parseError error
		page, parseError = strconv.ParseInt(stringified, 10, 64)
		if parseError != nil {
			traces.Trace(connection.ErrorLog, parseError)
			return 0
		}
	}

	// Guard against unusual pages.
	if page <= 0 {
		page = 0
	}
	return page
}

func Board(connection *connections.Connection) {
	// Find page.
	page := ReceivePage(connection)

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
		connection.SendView(views.View{Name: "Board", Data: map[string]any{
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

	session := sessions.New(connection, state.State{}).Start()
	defer session.Save()

	// Send the views.
	connection.SendView(views.View{Name: "Board", Data: map[string]any{
		"verified": session.State.Verified,
		"expired":  session.State.Expired,
		"page":     page,
		"hasMore":  hasMore,
		"articles": articles,
	}})
}
