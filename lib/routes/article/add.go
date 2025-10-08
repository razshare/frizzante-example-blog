package article

import (
	"context"
	"main/lib/core/clients"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/database"
	"main/lib/database/sqlc"
	"main/lib/sessions"
	"strings"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

func Add(client *clients.Client) {
	var title string
	var content string
	var id *uuid.UUID
	var err error

	if title = strings.Trim(receive.FormValue(client, "title"), " "); title == "" {
		send.Navigate(client, "/form?error=article title cannot be empty")
		return
	}

	if content = strings.Trim(receive.FormValue(client, "content"), " "); content == "" {
		send.Navigate(client, "/form?error=article content cannot be empty")
		return
	}

	state := sessions.Start(receive.SessionId(client))

	if id, err = uuid.NewV4(); nil != err {
		send.Navigatef(client, "/form?error=%s", err)
		return
	}

	if err = database.Queries.AddArticle(context.Background(), sqlc.AddArticleParams{
		ID:        id.String(),
		Title:     title,
		Content:   content,
		AccountID: state.AccountId,
		CreatedAt: time.Now().Unix(),
	}); err != nil {
		send.Navigatef(client, "/form?error=%s", err)
		return
	}

	send.Navigate(client, "/board")
}
