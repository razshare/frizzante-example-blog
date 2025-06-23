package handlers

import (
	"context"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/razshare/frizzante/libcon"
	"github.com/razshare/frizzante/libsession"
	"github.com/razshare/frizzante/libview"

	"main/lib"
	"main/lib/database"
	"main/lib/utilities/sqlc"
	"strings"
	"time"
)

func ArticleFormAction(con *libcon.Connection) {
	// Find page.
	form := con.ReceiveForm()
	title := strings.Trim(form.Get("title"), " ")
	content := strings.Trim(form.Get("content"), " ")

	if title == "" {
		con.SendView(libview.View{Name: "ArticleForm", Data: map[string]any{
			"error": "article title cannot be empty",
		}})
		return
	}

	if content == "" {
		con.SendView(libview.View{Name: "ArticleForm", Data: map[string]any{
			"error": "article content cannot be empty",
		}})
		return
	}

	state, _ := libsession.Session(con, lib.State{})
	articleId, articleIdError := uuid.NewV4()
	if nil != articleIdError {
		con.SendView(libview.View{Name: "Board", Data: map[string]any{
			"error": articleIdError.Error(),
		}})
		return
	}

	addArticleError := database.Queries.AddArticle(context.Background(), sqlc.AddArticleParams{
		ID:        articleId.String(),
		Title:     title,
		AccountID: state.AccountId,
		CreatedAt: time.Now().Unix(),
	})

	if nil != addArticleError {
		con.SendView(libview.View{Name: "Board", Data: map[string]any{
			"error": addArticleError.Error(),
		}})
		return
	}

	articleContentId, articleContentIdError := uuid.NewV4()
	if nil != articleContentIdError {
		con.SendView(libview.View{Name: "Board", Data: map[string]any{
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
		con.SendView(libview.View{Name: "Board", Data: map[string]any{
			"error": addContentError.Error(),
		}})
		return
	}

	con.SendNavigate("/board")
}
