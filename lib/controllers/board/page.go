package board

import (
	f "github.com/razshare/frizzante"
	"main/lib/config"
)

var guards = []f.Guard{
	config.GuardExpired,
	config.GuardVerified,
}

func init() {
	config.Server.LoadController(func(controller *f.Controller) {
		controller.
			WithBase(guards, base).
			WithAction(guards, base)
	})
}

type data struct {
	Articles []article `json:"articles"`
}

type article struct {
	ArticleId string `json:"articleId"`
	Title     string `json:"title"`
	CreatedAt int    `json:"createdAt"`
	AccountId string `json:"accountId"`
}

func base(req *f.Request, res *f.Response) {
	fetchNextArticle, closeFetch := config.FindArticles(0, 10)
	defer closeFetch()

	var articleId string
	var title string
	var createdAt int
	var accountId string
	var articles []article
	for fetchNextArticle(&articleId, &title, &createdAt, &accountId) {
		articles = append(articles, article{
			AccountId: accountId,
			Title:     title,
			CreatedAt: createdAt,
			ArticleId: articleId,
		})
	}

	res.SendView(f.NewViewWithData(f.RenderModeFull, data{Articles: articles}))
}
