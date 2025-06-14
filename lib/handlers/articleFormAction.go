package handlers

import (
	"context"
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/razshare/frizzante/frz"
	"main/lib"
	"main/lib/database"
	"main/lib/generated"
	"time"
)

func ArticleFormAction(c *frz.Connection) {
	// Find page.
	page := c.ReceiveQuery("page")

	state, _ := frz.Session(c, lib.State{})
	articleId, articleIdError := uuid.NewV4()
	if nil != articleIdError {
		c.SendView(frz.View{Name: "Board", Data: map[string]any{
			"error": articleIdError.Error(),
		}})
		return
	}

	form := c.ReceiveForm()
	addArticleError := database.Queries.AddArticle(context.Background(), generated.AddArticleParams{
		ID:        articleId.String(),
		Title:     form.Get("title"),
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

	addContentError := database.Queries.AddArticleContent(context.Background(), generated.AddArticleContentParams{
		ID:        articleContentId.String(),
		ArticleID: articleId.String(),
		Content:   form.Get("content"),
	})

	if nil != addContentError {
		c.SendView(frz.View{Name: "Board", Data: map[string]any{
			"error": addContentError.Error(),
		}})
		return
	}

	if page != "" {
		c.SendNavigate(fmt.Sprintf("/board?page=%s", page))
		return
	}

	c.SendNavigate("/board")
}
