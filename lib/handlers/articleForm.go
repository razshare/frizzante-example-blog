package handlers

import (
	"context"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib"
	"main/lib/database"
	"main/lib/database/sqlc"
	"strings"
	"time"
)

func ArticleForm(con *connections.Connection) {
	con.SendView(views.View{Name: "ArticleForm"})
}

func ArticleFormAction(con *connections.Connection) {
	// Find page.
	form := con.ReceiveForm()
	title := strings.Trim(form.Get("title"), " ")
	content := strings.Trim(form.Get("content"), " ")

	if title == "" {
		con.SendView(views.View{Name: "ArticleForm", Data: map[string]any{
			"error": "article title cannot be empty",
		}})
		return
	}

	if content == "" {
		con.SendView(views.View{Name: "ArticleForm", Data: map[string]any{
			"error": "article content cannot be empty",
		}})
		return
	}

	session := sessions.New(con, lib.State{}).Start()
	defer session.Save()

	articleId, articleIdError := uuid.NewV4()
	if nil != articleIdError {
		con.SendView(views.View{Name: "Board", Data: map[string]any{
			"error": articleIdError.Error(),
		}})
		return
	}

	addArticleError := database.Queries.AddArticle(context.Background(), sqlc.AddArticleParams{
		ID:        articleId.String(),
		Title:     title,
		AccountID: session.State.AccountId,
		CreatedAt: time.Now().Unix(),
	})

	if nil != addArticleError {
		con.SendView(views.View{Name: "Board", Data: map[string]any{
			"error": addArticleError.Error(),
		}})
		return
	}

	articleContentId, articleContentIdError := uuid.NewV4()
	if nil != articleContentIdError {
		con.SendView(views.View{Name: "Board", Data: map[string]any{
			"error": articleContentIdError.Error(),
		}})
		return
	}

	addContentError := database.Queries.AddArticleContent(context.Background(), sqlc.AddArticleContentParams{
		ID:        articleContentId.String(),
		ArticleID: articleId.String(),
		Content:   content,
	})

	if nil != addContentError {
		con.SendView(views.View{Name: "Board", Data: map[string]any{
			"error": addContentError.Error(),
		}})
		return
	}

	con.SendNavigate("/board")
}
