package routes

import (
	"context"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/razshare/frizzante"
	"main/lib/database"
	"main/lib/generated"
	"main/lib/guards"
	"main/lib/sessions"
	"main/lib/value"
	"time"
)

func GetBoard(req *frizzante.Request, res *frizzante.Response) {
	if !frizzante.AllGuardsPass(req, res, guards.NotExpired, guards.Verified) {
		return
	}

	articles := value.Wrap(database.Queries.SqlFindArticles(context.Background(), generated.SqlFindArticlesParams{
		Offset: 0,
		Limit:  10,
	}))

	if !articles.Ok() {
		res.SendView(frizzante.View{Name: "Board", Error: articles.Error})
		return
	}

	res.SendView(frizzante.View{Name: "Board", Data: articles.Value})
}

func PostBoard(req *frizzante.Request, res *frizzante.Response) {
	if !frizzante.AllGuardsPass(req, res, guards.NotExpired, guards.Verified) {
		return
	}

	articleId := value.Wrap(uuid.NewV4())
	if !articleId.Ok() {
		res.SendView(frizzante.View{Name: "Board", Error: articleId.Error})
		return
	}

	form := req.ReceiveForm()
	session := frizzante.SessionStart(req, res, sessions.Adapter)
	addArticle := value.WrapNothing(database.Queries.SqlAddArticle(context.Background(), generated.SqlAddArticleParams{
		ID:        articleId.Value.String(),
		AccountID: session.Data.AccountId,
		CreatedAt: int32(time.Now().Unix()),
	}))
	if !addArticle.Ok() {
		res.SendView(frizzante.View{Name: "Board", Error: addArticle.Error})
		return
	}

	articleContentId := value.Wrap(uuid.NewV4())
	if !articleContentId.Ok() {
		res.SendView(frizzante.View{Name: "Board", Error: articleContentId.Error})
		return
	}

	addContent := value.WrapNothing(database.Queries.SqlAddArticleContent(context.Background(), generated.SqlAddArticleContentParams{
		ID:        articleContentId.Value.String(),
		ArticleID: articleId.Value.String(),
		Title:     form.Get("title"),
		Content:   form.Get("content"),
	}))

	if !addContent.Ok() {
		res.SendView(frizzante.View{Name: "Board", Error: addContent.Error})
		return
	}

	res.SendNavigate("/board")
}
