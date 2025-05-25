package lib

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	f "github.com/razshare/frizzante"
	"main/lib/sqlc"
	"time"
)

func init() {
	Server.
		WithRequestHandler("GET /account", func(req *f.Request, res *f.Response) {
			if !f.AllGuardsPass(req, res, NotExpired, Verified) {
				return
			}

			session := f.SessionStart(req, res, SessionAdapter)
			account, err := Queries.SqlFindAccountById(context.Background(), session.Data.AccountId)

			if nil != err {
				res.SendView(NewErrorView(err))
				return
			}

			res.SendView(f.View{
				Name: "Account",
				Data: account,
			})
		}).
		WithRequestHandler("GET /board", func(req *f.Request, res *f.Response) {
			if !f.AllGuardsPass(req, res, NotExpired, Verified) {
				return
			}

			articles, err := Queries.SqlFindArticles(context.Background(), sqlc.SqlFindArticlesParams{
				Offset: 0,
				Limit:  10,
			})

			if err != nil {
				res.SendView(NewErrorView(err))
				return
			}

			res.SendView(f.View{
				Name: "Board",
				Data: map[string]any{
					"articles": articles,
				},
			})
		}).
		WithRequestHandler("GET /expired", func(req *f.Request, res *f.Response) {
			res.SendView(f.View{
				Name: "Expired",
			})
		}).
		WithRequestHandler("GET /logout", func(req *f.Request, res *f.Response) {
			session := f.SessionStart(req, res, SessionAdapter)
			session.Data.Verified = false
			session.Save()
			res.SendNavigate("/login")
		}).
		WithRequestHandler("POST /logout", func(req *f.Request, res *f.Response) {
			session := f.SessionStart(req, res, SessionAdapter)
			session.Data.Verified = false
			session.Save()
			res.SendNavigate("/login")
		}).
		WithRequestHandler("GET /login", func(req *f.Request, res *f.Response) {
			res.SendView(f.View{
				Name: "Login",
			})
		}).
		WithRequestHandler("POST /login", func(req *f.Request, res *f.Response) {
			form := req.ReceiveForm()
			id := form.Get("id")
			password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

			_, err := Queries.SqlVerifyAccount(context.Background(), sqlc.SqlVerifyAccountParams{
				ID:       id,
				Password: password,
			})

			if err != nil {
				res.SendView(NewErrorView(err))
				return
			}

			session := f.SessionStart(req, res, SessionAdapter)
			session.Data.Verified = true
			session.Data.AccountId = id
			session.Save()
			res.SendNavigate("/board")
		}).
		WithRequestHandler("GET /register", func(req *f.Request, res *f.Response) {
			res.SendView(f.View{
				Name: "Register",
			})
		}).
		WithRequestHandler("POST /register", func(req *f.Request, res *f.Response) {
			form := req.ReceiveForm()
			accountId := form.Get("id")

			_, err := Queries.SqlFindAccountById(context.Background(), accountId)
			if err != nil {
				return
			}

			if nil != err {
				res.SendView(NewErrorView(err))
				return
			}

			accountDisplayName := form.Get("displayName")
			rawPassword := form.Get("password")

			if "" == accountId || "" == accountDisplayName || "" == rawPassword {
				res.SendView(NewErrorView(errors.New("please fill all fields")))
				return
			}

			accountPassword := fmt.Sprintf("%x", sha256.Sum256([]byte(rawPassword)))
			err = Queries.SqlAddAccount(context.Background(), sqlc.SqlAddAccountParams{
				ID:          accountId,
				DisplayName: accountDisplayName,
				Password:    accountPassword,
			})
			if err != nil {
				res.SendView(NewErrorView(err))
				return
			}
			res.SendNavigate("/login")
		}).
		WithRequestHandler("POST /article", func(req *f.Request, res *f.Response) {
			if !f.AllGuardsPass(req, res, NotExpired, Verified) {
				return
			}

			articleId, err := uuid.NewV4()
			if err != nil {
				res.SendView(NewErrorView(err))
				return
			}

			form := req.ReceiveForm()
			session := f.SessionStart(req, res, SessionAdapter)
			err = Queries.SqlAddArticle(context.Background(), sqlc.SqlAddArticleParams{
				ID:        articleId.String(),
				AccountID: session.Data.AccountId,
				CreatedAt: int32(time.Now().Unix()),
			})
			if err != nil {
				res.SendView(NewErrorView(err))
				return
			}

			articleContentId, err := uuid.NewV4()
			if err != nil {
				res.SendView(NewErrorView(err))
				return
			}

			err = Queries.SqlAddArticleContent(context.Background(), sqlc.SqlAddArticleContentParams{
				ID:        articleContentId.String(),
				ArticleID: articleId.String(),
				Title:     form.Get("title"),
				Content:   form.Get("content"),
			})

			if err != nil {
				res.SendView(NewErrorView(err))
				return
			}

			res.SendNavigate("/board")
		}).
		WithRequestHandler("GET /", func(req *f.Request, res *f.Response) {
			res.SendFileOrElse(func() {
				res.SendView(f.View{
					Name: "Login",
				})
			})
		})
}
