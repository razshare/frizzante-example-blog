package pages

import (
	f "github.com/razshare/frizzante"
	"main/lib/guards"
	"main/lib/sql"
)

type BoardArticle struct {
	ArticleId string `json:"articleId"`
	Title     string `json:"title"`
	CreatedAt int    `json:"createdAt"`
	AccountId string `json:"accountId"`
}

func Board(page *f.Page) {
	f.PageWithPath(page, "/board")
	f.PageWithView(page, f.ViewReference("Board"))
	f.PageWithGuardHandler(page, guards.Verified)
	f.PageWithBaseHandler(page, func(request *f.Request, response *f.Response, view *f.View) {
		fetchNextArticle, closeFetch := sql.FindArticles(0, 10)
		defer closeFetch()

		var articleId string
		var title string
		var createdAt int
		var accountId string
		var articles []BoardArticle
		for fetchNextArticle(&articleId, &title, &createdAt, &accountId) {
			articles = append(articles, BoardArticle{
				AccountId: accountId,
				Title:     title,
				CreatedAt: createdAt,
				ArticleId: articleId,
			})
		}

		f.ViewWithData(view, "articles", articles)
	})
	f.PageWithActionHandler(page, func(_ *f.Request, response *f.Response, _ *f.View) {
		f.ResponseSendNavigate(response, "Login")
	})
}
