package handlers

import (
	"context"
	"github.com/razshare/frizzante/libcon"
	"github.com/razshare/frizzante/libsession"
	"github.com/razshare/frizzante/libview"

	"main/lib"
	"main/lib/database"
	"main/lib/notifiers"
	"main/lib/utilities/sqlc"
	"strconv"
)

var pageSize int64 = 10

func ReceivePage(con *libcon.Connection) int64 {
	var page int64

	// Find page.
	if stringified := con.ReceiveQuery("page"); stringified != "" {
		var parseError error
		page, parseError = strconv.ParseInt(stringified, 10, 64)
		if parseError != nil {
			notifiers.Console.SendError(parseError)
			return 0
		}
	}

	// Guard against unusual pages.
	if page <= 0 {
		page = 0
	}
	return page
}

func Board(con *libcon.Connection) {
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
		con.SendView(libview.View{Name: "Board", Data: map[string]any{
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

	state, _ := libsession.Session(con, lib.State{})

	// Send the view.
	con.SendView(libview.View{Name: "Board", Data: map[string]any{
		"verified": state.Verified,
		"expired":  state.Expired,
		"page":     page,
		"hasMore":  hasMore,
		"articles": articles,
	}})
}
