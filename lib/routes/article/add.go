package article

import (
	"context"
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/core/view"
	"main/lib/database"
	"main/lib/database/sqlc"
	"main/lib/session"
	"strings"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

func Add(c *client.Client) {
	f := receive.Form(c)
	t := strings.Trim(f.Get("title"), " ")
	cn := strings.Trim(f.Get("content"), " ")

	if t == "" {
		send.View(c, view.View{Name: "ArticleForm", Props: map[string]any{
			"error": "article title cannot be empty",
		}})
		return
	}

	if cn == "" {
		send.View(c, view.View{Name: "ArticleForm", Props: map[string]any{
			"error": "article content cannot be empty",
		}})
		return
	}

	s := session.Start(receive.SessionId(c))

	aid, err := uuid.NewV4()
	if nil != err {
		send.View(c, view.View{Name: "ArticleForm", Props: map[string]any{
			"error": err.Error(),
		}})
		return
	}

	err = database.Queries.AddArticle(context.Background(), sqlc.AddArticleParams{
		ID:        aid.String(),
		Title:     t,
		Content:   cn,
		AccountID: s.AccountId,
		CreatedAt: time.Now().Unix(),
	})

	if err != nil {
		send.View(c, view.View{Name: "ArticleForm", Props: map[string]any{
			"error": err.Error(),
		}})
		return
	}

	send.Navigate(c, "/board")
}
