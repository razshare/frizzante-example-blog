package handlers

import (
	"context"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/razshare/frizzante/frz"
	"main/lib"
	"main/lib/database"
	"main/lib/utilities/sqlc"
	"strings"
	"time"
)

func ArticleFormAction(c *frz.Connection) {
	// Find page.
	form := c.ReceiveForm()
	title := strings.Trim(form.Get("title"), " ")
	content := strings.Trim(form.Get("content"), " ")

	if title == "" {
		c.SendView(frz.View{Name: "ArticleForm", Data: map[string]any{
			"error": "article title cannot be empty",
		}})
		return
	}

	if content == "" {
		c.SendView(frz.View{Name: "ArticleForm", Data: map[string]any{
			"error": "article content cannot be empty",
		}})
		return
	}

	state, _ := frz.Session(c, lib.State{})
	articleId, articleIdError := uuid.NewV4()
	if nil != articleIdError {
		c.SendView(frz.View{Name: "Board", Data: map[string]any{
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
		c.SendView(frz.View{Name: "Board", Data: map[string]any{
			"error": addArticleError.Error(),
		}})
		return
	}

	articleContentId, articleContentIdError := uuid.NewV4()
	if nil != articleContentIdError {
		c.SendView(frz.View{Name: "Board", Data: map[string]any{
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
		c.SendView(frz.View{Name: "Board", Data: map[string]any{
			"error": addContentError.Error(),
		}})
		return
	}

	c.SendNavigate("/board")
}
