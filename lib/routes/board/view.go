package board

import (
	"context"
	"main/lib/core/clients"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/core/types"
	"main/lib/core/views"
	"main/lib/database"
	"main/lib/database/sqlc"
	"main/lib/sessions"
)

func View(client *clients.Client) {
	var page = Paginate(client)
	var articles []sqlc.Article
	var count int
	var hasMore bool
	var err error

	session := sessions.Start(receive.SessionId(client))

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
		send.View(client, views.View{
			Name: "Board",
			Props: map[string]any{
				"error":    err.Error(),
				"loggedIn": session.LoggedIn,
				"expired":  session.LoginExpired,
			},
		})
		return
	}

	if articles == nil {
		articles = make([]sqlc.Article, 0)
	}

	count = len(articles)
	if hasMore = count == int(PageSize)+1; hasMore {
		if articles = articles[:count-1]; articles == nil {
			articles = make([]sqlc.Article, 0)
		}
	}

	// Send the views.
	send.View(client, views.View{Name: "Board", Props: Props{
		Page:     page,
		HasMore:  hasMore,
		Articles: articles,
		LoggedIn: session.LoggedIn,
		Expired:  session.LoginExpired,
		Error:    receive.Query(client, "error"),
	}})
}

type Props struct {
	Page     int64          `json:"page"`
	HasMore  bool           `json:"hasMore"`
	Articles []sqlc.Article `json:"articles"`
	LoggedIn bool           `json:"loggedIn"`
	Expired  bool           `json:"expired"`
	Error    string         `json:"error"`
}

func init() {
	_ = types.Generate[Props]()
}
