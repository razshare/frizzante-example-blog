package handlers

import (
	"context"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/razshare/frizzante"
	"main/lib"
	"main/lib/database"
	"main/lib/generated"
	"time"
)

func GetBoard(c *frizzante.Connection) {
	lib.SessionStartProtected(c, func(state *lib.State) {
		articles, articleError := database.Queries.SqlFindArticles(context.Background(), generated.SqlFindArticlesParams{
			Offset: 0,
			Limit:  10,
		})

		if nil != articleError {
			c.SendView(frizzante.View{Name: "Board", Error: articleError})
			return
		}

		c.SendView(frizzante.View{Name: "Board", Data: articles})
	})
}

func PostBoard(c *frizzante.Connection) {
	lib.SessionStartProtected(c, func(state *lib.State) {
		articleId, articleIdError := uuid.NewV4()
		if nil != articleIdError {
			c.SendView(frizzante.View{Name: "Board", Error: articleIdError})
			return
		}

		form := c.ReceiveForm()
		addArticleError := database.Queries.SqlAddArticle(context.Background(), generated.SqlAddArticleParams{
			ID:        articleId.String(),
			AccountID: state.AccountId,
			CreatedAt: time.Now().Unix(),
		})
		if nil != addArticleError {
			c.SendView(frizzante.View{Name: "Board", Error: addArticleError})
			return
		}

		articleContentId, articleContentIdError := uuid.NewV4()
		if nil != articleContentIdError {
			c.SendView(frizzante.View{Name: "Board", Error: articleContentIdError})
			return
		}

		addContentError := database.Queries.SqlAddArticleContent(context.Background(), generated.SqlAddArticleContentParams{
			ID:        articleContentId.String(),
			ArticleID: articleId.String(),
			Title:     form.Get("title"),
			Content:   form.Get("content"),
		})

		if nil != addContentError {
			c.SendView(frizzante.View{Name: "Board", Error: addContentError})
			return
		}

		c.SendNavigate("/board")
	})
}
