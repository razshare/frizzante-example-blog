package article

import (
	"context"
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/database"
	"main/lib/database/sqlc"
	"main/lib/session/memory"
	"strings"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

func Add(client *client.Client) {
	var title string
	var content string
	var id *uuid.UUID
	var err error

	if title = strings.Trim(receive.FormValue(client, "title"), " "); title == "" {
		send.Navigate(client, "/article-form?error=article title cannot be empty")
		return
	}

	if content = strings.Trim(receive.FormValue(client, "content"), " "); content == "" {
		send.Navigate(client, "/article-form?error=article content cannot be empty")
		return
	}

	state := memory.Start(receive.SessionId(client))

	if id, err = uuid.NewV4(); nil != err {
		send.Navigatef(client, "/article-form?error=%s", err.Error())
		return
	}

	if err = database.Queries.AddArticle(context.Background(), sqlc.AddArticleParams{
		ID:        id.String(),
		Title:     title,
		Content:   content,
		AccountID: state.AccountId,
		CreatedAt: time.Now().Unix(),
	}); err != nil {
		send.Navigatef(client, "/article-form?error=%s", err.Error())
		return
	}

	send.Navigate(client, "/board")
}
