package pages

import (
	"crypto/sha256"
	"fmt"
	f "github.com/razshare/frizzante"
	"main/lib/sql"
)

func Register(page *f.Page) {
	f.PageWithPath(page, "/register")
	f.PageWithView(page, f.ViewReference("Register"))
	f.PageWithActionHandler(page, func(request *f.Request, response *f.Response, view *f.View) {
		form := f.RequestReceiveForm(request)
		id := form.Get("id")
		if sql.AccountExists(id) {
			f.ViewWithData(view, "error", fmt.Sprintf("Account %s already exists.", id))
			return
		}

		displayName := form.Get("displayName")
		rawPassword := form.Get("password")

		if "" == id || "" == displayName || "" == rawPassword {
			f.ViewWithData(view, "error", "Please fill all fields.")
			return
		}

		password := fmt.Sprintf("%x", sha256.Sum256([]byte(rawPassword)))
		sql.AddAccount(id, displayName, password)
		f.ResponseSendNavigate(response, "Login")
	})
}
