package pages

import (
	f "github.com/razshare/frizzante"
	"main/lib"
)

var Board = f.NewPageController().WithBase(boardBase)

type BoardArticle struct {
	ArticleId string `json:"articleId"`
	Title     string `json:"title"`
	CreatedAt int    `json:"createdAt"`
	AccountId string `json:"accountId"`
}

type BoardData struct {
	Articles []BoardArticle `json:"articles"`
}

func boardBase(req *f.Request, res *f.Response) {
	fetchNextArticle, closeFetch := lib.FindArticles(0, 10)
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

	res.SendView(f.NewViewWithData(f.RenderModeFull, BoardData{Articles: articles}))
}
