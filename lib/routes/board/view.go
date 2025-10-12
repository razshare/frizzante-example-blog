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

func init() {
	_ = types.Generate[Props]()
}

type Props struct {
	Page     int64          `json:"page"`
	HasMore  bool           `json:"hasMore"`
	Articles []sqlc.Article `json:"articles"`
	LoggedIn bool           `json:"loggedIn"`
	Expired  bool           `json:"expired"`
	Error    string         `json:"error"`
}

func View(client *clients.Client) {
	session := sessions.Start(receive.SessionId(client))
	page := Paginate(client)

	var err error
	var articles []sqlc.Article

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
		send.View(client, views.View{Name: "Board", Props: Props{
			Error:    err.Error(),
			LoggedIn: session.LoggedIn,
			Expired:  session.LoginExpired,
		}})
		return
	}

	if articles == nil {
		articles = make([]sqlc.Article, 0)
	}

	count := len(articles)
	hasMore := count == int(PageSize)+1

	if hasMore {
		if articles = articles[:count-1]; articles == nil {
			articles = make([]sqlc.Article, 0)
		}
	}

	send.View(client, views.View{Name: "Board", Props: Props{
		Page:     page,
		HasMore:  hasMore,
		Articles: articles,
		LoggedIn: session.LoggedIn,
		Expired:  session.LoginExpired,
		Error:    session.BoardError,
	}})
}
