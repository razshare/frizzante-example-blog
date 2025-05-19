package pages

import (
	f "github.com/razshare/frizzante"
	"main/lib"
)

type BoardArticle struct {
	ArticleId string `json:"articleId"`
	Title     string `json:"title"`
	CreatedAt int    `json:"createdAt"`
	AccountId string `json:"accountId"`
}

type BoardData struct {
	Articles []BoardArticle `json:"articles"`
}

type BoardController struct {
	f.PageController
}

func (_ BoardController) Configure() f.PageConfiguration {
	return f.PageConfiguration{
		Path: "/board",
	}
}

func (_ BoardController) Base(request *f.Request, response *f.Response) {
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

	response.SendView(f.NewView(BoardData{Articles: articles}))
}

func (_ BoardController) Action(_ *f.Request, response *f.Response) {
	response.SendView(f.NewView(false))
}
