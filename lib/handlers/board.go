package handlers

import (
	"context"
	"github.com/razshare/frizzante/frz"
	"main/lib"
	"main/lib/database"
	"main/lib/notifiers"
	"main/lib/utilities/sqlc"
	"strconv"
)

var pageSize int64 = 10

func ReceivePage(c *frz.Connection) int64 {
	var p int64

	// Find page.
	if stringified := c.ReceiveQuery("page"); stringified != "" {
		var parseError error
		p, parseError = strconv.ParseInt(stringified, 10, 64)
		if parseError != nil {
			notifiers.Console.SendError(parseError)
			return 0
		}
	}

	// Guard against unusual pages.
	if p <= 0 {
		p = 0
	}
	return p
}

func Board(c *frz.Connection) {
	// Find page.
	page := ReceivePage(c)

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
		c.SendView(frz.View{Name: "Board", Data: map[string]any{
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

	state, _ := frz.Session(c, lib.State{})

	// Send the view.
	c.SendView(frz.View{Name: "Board", Data: map[string]any{
		"verified": state.Verified,
		"expired":  state.Expired,
		"page":     page,
		"hasMore":  hasMore,
		"articles": articles,
	}})
}
