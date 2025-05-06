package pages

import (
	f "github.com/razshare/frizzante"
	"main/lib/sql"
)

type BoardArticle struct {
	ArticleId string `json:"articleId"`
	Title     string `json:"title"`
	CreatedAt int    `json:"createdAt"`
	AccountId string `json:"accountId"`
}

func Board(context f.PageContext) {
	// Context.
	path, view, base, action := context()

	// Configure.
	path("/board")
	view(f.ViewReference("Board"))
	base(func(request *f.Request, response *f.Response, view *f.View) {
		get, _, _ := f.SessionStart(request, response)
		if !get("verified", false).(bool) {
			f.ResponseSendNavigate(response, "Login")
		}

		article, closeFetch := sql.FindArticles(0, 10)
		defer closeFetch()

		var articleId string
		var title string
		var createdAt int
		var accountId string

		var articles []BoardArticle
		for article(&articleId, &title, &createdAt, &accountId) {
			articles = append(articles, BoardArticle{
				AccountId: accountId,
				Title:     title,
				CreatedAt: createdAt,
				ArticleId: articleId,
			})
		}

		f.ViewWithData(view, "articles", articles)
	})
	action(func(_ *f.Request, response *f.Response, _ *f.View) {
		f.ResponseSendNavigate(response, "Login")
	})
}
