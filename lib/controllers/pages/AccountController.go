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

type AccountController struct {
	f.PageController
}

func (_ AccountController) Configure() f.PageConfiguration {
	return f.PageConfiguration{
		Path: "/account",
	}
}

func (_ AccountController) Base(request *f.Request, response *f.Response) {
	session := f.SessionStart(request, response, sessions.Archived)

	fetchAccount, closeFetch := lib.FindAccountById(session.Data.AccountId)
	defer closeFetch()

	var accountId string
	var displayName string
	var createdAt int
	var updatedAt int

	fetchAccount(&accountId, &displayName, &createdAt, &updatedAt)

	response.SendView(f.NewViewWithData(f.RenderModeFull, AccountData{
		AccountId:   accountId,
		DisplayName: displayName,
	}))
	session.Save()
}

func (_ AccountController) Action(request *f.Request, response *f.Response) {
	response.SendView(f.NewViewWithData(f.RenderModeFull, AccountData{}))
}
