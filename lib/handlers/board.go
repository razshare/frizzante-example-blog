package handlers

import (
	"context"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/razshare/frizzante/frz"
	"main/lib"
	"main/lib/database"
	"main/lib/generated"
	"time"
)

func GetBoard(c *frz.Connection) {
	articles, articleError := database.Queries.SqlFindArticles(
		context.Background(),
		generated.SqlFindArticlesParams{
			Offset: 0,
			Limit:  10,
		},
	)

	if nil != articleError {
		c.SendView(frz.View{Name: "Board", Error: articleError.Error()})
		return
	}

	c.SendView(frz.View{Name: "Board", Data: articles})
}

func PostBoard(c *frz.Connection) {
	state, _ := frz.Session(c, lib.State{})
	articleId, articleIdError := uuid.NewV4()
	if nil != articleIdError {
		c.SendView(frz.View{Name: "Board", Error: articleIdError.Error()})
		return
	}

	form := c.ReceiveForm()
	addArticleError := database.Queries.SqlAddArticle(context.Background(), generated.SqlAddArticleParams{
		ID:        articleId.String(),
		AccountID: state.AccountId,
		CreatedAt: time.Now().Unix(),
	})
	if nil != addArticleError {
		c.SendView(frz.View{Name: "Board", Error: addArticleError.Error()})
		return
	}

	articleContentId, articleContentIdError := uuid.NewV4()
	if nil != articleContentIdError {
		c.SendView(frz.View{Name: "Board", Error: articleContentIdError.Error()})
		return
	}

	addContentError := database.Queries.SqlAddArticleContent(context.Background(), generated.SqlAddArticleContentParams{
		ID:        articleContentId.String(),
		ArticleID: articleId.String(),
		Title:     form.Get("title"),
		Content:   form.Get("content"),
	})

	if nil != addContentError {
		c.SendView(frz.View{Name: "Board", Error: addContentError.Error()})
		return
	}

	c.SendNavigate("/board")
}
