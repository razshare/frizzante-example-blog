package article_form

import (
	"context"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/razshare/frizzante/client"
	"github.com/razshare/frizzante/receive"
	"github.com/razshare/frizzante/send"
	"github.com/razshare/frizzante/view"
	"main/lib/database"
	"main/lib/database/sqlc"
	"main/lib/session"
	"strings"
	"time"
)

func Action(c *client.Client) {
	f := receive.Form(c)
	t := strings.Trim(f.Get("title"), " ")
	cn := strings.Trim(f.Get("content"), " ")

	if t == "" {
		send.View(c, view.View{Name: "ArticleForm", Data: map[string]any{
			"error": "article title cannot be empty",
		}})
		return
	}

	if cn == "" {
		send.View(c, view.View{Name: "ArticleForm", Data: map[string]any{
			"error": "article content cannot be empty",
		}})
		return
	}

	s := session.Start(receive.SessionId(c))

	aid, err := uuid.NewV4()
	if nil != err {
		send.View(c, view.View{Name: "Board", Data: map[string]any{
			"error": err.Error(),
		}})
		return
	}

	err = database.Queries.AddArticle(context.Background(), sqlc.AddArticleParams{
		ID:        aid.String(),
		Title:     t,
		AccountID: s.AccountId,
		CreatedAt: time.Now().Unix(),
	})

	if err != nil {
		send.View(c, view.View{Name: "Board", Data: map[string]any{
			"error": err.Error(),
		}})
		return
	}

	acid, err := uuid.NewV4()
	if err != nil {
		send.View(c, view.View{Name: "Board", Data: map[string]any{
			"error": err.Error(),
		}})
		return
	}

	err = database.Queries.AddArticleContent(context.Background(), sqlc.AddArticleContentParams{
		ID:        acid.String(),
		ArticleID: aid.String(),
		Content:   cn,
	})

	if err != nil {
		send.View(c, view.View{Name: "Board", Data: map[string]any{
			"error": err.Error(),
		}})
		return
	}

	send.Navigate(c, "/board")
}
