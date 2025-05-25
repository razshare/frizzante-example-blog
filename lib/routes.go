package lib

import (
	"crypto/sha256"
	"fmt"
	f "github.com/razshare/frizzante"
)

func init() {
	Server.
		WithRequestHandler("GET /account", func(req *f.Request, res *f.Response) {
			session := f.SessionStart(req, res, SessionAdapter)

			fetchAccount, closeFetch := FindAccountById(session.Data.AccountId)
			defer closeFetch()

			var accountId string
			var displayName string
			var createdAt int
			var updatedAt int

			fetchAccount(&accountId, &displayName, &createdAt, &updatedAt)

			res.SendView(f.View{
				Name: "Account",
				Data: map[string]string{
					"accountId":   accountId,
					"displayName": displayName,
				},
			})
		}).
		WithRequestHandler("GET /board", func(req *f.Request, res *f.Response) {
			fetchNextArticle, closeFetch := FindArticles(0, 10)
			defer closeFetch()

			var articleId string
			var title string
			var createdAt int
			var accountId string
			var articles []map[string]any
			for fetchNextArticle(&articleId, &title, &createdAt, &accountId) {
				articles = append(articles, map[string]any{
					"accountId": accountId,
					"title":     title,
					"createdAt": createdAt,
					"articleId": articleId,
				})
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

			if !VerifyAccount(id, password) {
				res.SendView(f.View{

					Name: "login",
					Data: map[string]any{
						"error": "Invalid credentials",
					},
				})
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
			id := form.Get("id")
			if AccountExists(id) {
				res.SendView(f.View{
					Name: "Register",
					Data: map[string]any{
						"error": fmt.Sprintf("Account %s already exists.", id),
					},
				})
				return
			}

			displayName := form.Get("displayName")
			rawPassword := form.Get("password")

			if "" == id || "" == displayName || "" == rawPassword {
				res.SendView(f.View{
					Name: "Register",
					Data: map[string]any{
						"error": "Please fill all fields.",
					},
				})
				return
			}

			password := fmt.Sprintf("%x", sha256.Sum256([]byte(rawPassword)))
			AddAccount(id, displayName, password)
			res.SendNavigate("/login")
		}).
		WithRequestHandler("GET /", func(req *f.Request, res *f.Response) {
			res.SendFileOrElse(func() {
				res.SendView(f.View{
					Name: "Welcome",
				})
			})
		})
}
