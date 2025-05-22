package pages

import (
	f "github.com/razshare/frizzante"
	"main/lib"
	"main/lib/sessions"
)

type AccountData struct {
	AccountId   string `json:"accountId"`
	DisplayName string `json:"displayName"`
}

var Account = f.NewPageController().WithBase(accountBase)

func accountBase(req *f.Request, res *f.Response) {
	session := f.SessionStart(req, res, sessions.Archived)

	fetchAccount, closeFetch := lib.FindAccountById(session.Data.AccountId)
	defer closeFetch()

	var accountId string
	var displayName string
	var createdAt int
	var updatedAt int

	fetchAccount(&accountId, &displayName, &createdAt, &updatedAt)

	res.SendView(f.NewViewWithData(f.RenderModeFull, AccountData{
		AccountId:   accountId,
		DisplayName: displayName,
	}))
	session.Save()
}
