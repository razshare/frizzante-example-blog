package indexes

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

func boardShowFunction(req *f.Request, res *f.Response, p *f.Page) {
	get, _, _ := f.SessionStart(req, res)
	if !get("verified", false).(bool) {
		f.SendNavigate(res, "login")
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

	f.PageWithData(p, "articles", articles)
}

func boardActionFunction(_ *f.Request, _ *f.Response, _ *f.Page) {
	// Noop.
}

func Board(
	route func(path string, page string),
	show func(showFunction f.PageFunction),
	action func(actionFunction f.PageFunction),
) {
	route("/board", "board")
	show(boardShowFunction)
	action(boardActionFunction)
}
