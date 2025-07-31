package handlers

import (
	"context"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib/database"
	"main/lib/database/sqlc"
	"main/lib/state"
	"strings"
	"time"
)

func ArticleForm(connection *connections.Connection) {
	connection.SendView(views.View{Name: "ArticleForm"})
}

func ArticleFormAction(connection *connections.Connection) {
	// Find page.
	form := connection.ReceiveForm()
	title := strings.Trim(form.Get("title"), " ")
	content := strings.Trim(form.Get("content"), " ")

	if title == "" {
		connection.SendView(views.View{Name: "ArticleForm", Data: map[string]any{
			"error": "article title cannot be empty",
		}})
		return
	}

	if content == "" {
		connection.SendView(views.View{Name: "ArticleForm", Data: map[string]any{
			"error": "article content cannot be empty",
		}})
		return
	}

	session := sessions.New(connection, state.State{}).Start()
	defer session.Save()

	articleId, articleIdError := uuid.NewV4()
	if nil != articleIdError {
		connection.SendView(views.View{Name: "Board", Data: map[string]any{
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
		connection.SendView(views.View{Name: "Board", Data: map[string]any{
			"error": addArticleError.Error(),
		}})
		return
	}

	articleContentId, articleContentIdError := uuid.NewV4()
	if nil != articleContentIdError {
		connection.SendView(views.View{Name: "Board", Data: map[string]any{
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
		connection.SendView(views.View{Name: "Board", Data: map[string]any{
			"error": addContentError.Error(),
		}})
		return
	}

	connection.SendNavigate("/board")
}
