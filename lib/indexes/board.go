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

func showAction(req *f.Request, res *f.Response, p *f.Page) {
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

func Board() (
	page string,
	show f.PageFunction,
	action f.PageFunction,
) {
	page = "board"
	show = showAction
	return
}
