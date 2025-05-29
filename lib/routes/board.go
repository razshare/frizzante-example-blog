package routes

import (
	"context"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/razshare/frizzante"
	"main/lib/database"
	"main/lib/generated"
	"main/lib/guards"
	"main/lib/sessions"
	"time"
)

func GetBoard(req *frizzante.Request, res *frizzante.Response) {
	if !frizzante.AllGuardsPass(req, res, guards.NotExpired, guards.Verified) {
		return
	}

	articles, articleError := database.Queries.SqlFindArticles(context.Background(), generated.SqlFindArticlesParams{
		Offset: 0,
		Limit:  10,
	})

	if nil != articleError {
		res.SendView(frizzante.View{Name: "Board", Error: articleError})
		return
	}

	res.SendView(frizzante.View{Name: "Board", Data: articles})
}

func PostBoard(req *frizzante.Request, res *frizzante.Response) {
	if !frizzante.AllGuardsPass(req, res, guards.NotExpired, guards.Verified) {
		return
	}

	articleId, articleIdError := uuid.NewV4()
	if nil != articleIdError {
		res.SendView(frizzante.View{Name: "Board", Error: articleIdError})
		return
	}

	form := req.ReceiveForm()
	session := frizzante.SessionStart(req, res, sessions.Adapter)

	addArticleError := database.Queries.SqlAddArticle(context.Background(), generated.SqlAddArticleParams{
		ID:        articleId.String(),
		AccountID: session.Data.AccountId,
		CreatedAt: time.Now().Unix(),
	})
	if nil != addArticleError {
		res.SendView(frizzante.View{Name: "Board", Error: addArticleError})
		return
	}

	articleContentId, articleContentIdError := uuid.NewV4()
	if nil != articleContentIdError {
		res.SendView(frizzante.View{Name: "Board", Error: articleContentIdError})
		return
	}

	addContentError := database.Queries.SqlAddArticleContent(context.Background(), generated.SqlAddArticleContentParams{
		ID:        articleContentId.String(),
		ArticleID: articleId.String(),
		Title:     form.Get("title"),
		Content:   form.Get("content"),
	})

	if nil != addContentError {
		res.SendView(frizzante.View{Name: "Board", Error: addContentError})
		return
	}

	res.SendNavigate("/board")
}
