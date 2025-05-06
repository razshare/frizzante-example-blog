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

func boardShowFunction(request *f.Request, response *f.Response, view *f.View) {
	get, _, _ := f.SessionStart(request, response)
	if !get("verified", false).(bool) {
		f.SendNavigate(response, "login")
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
}

func Board(context f.PageContext) {
	path, view, base, _ := context()
	path("/board")
	view(f.ViewReference("Board"))
	base(boardShowFunction)
}
