package account

import (
	f "github.com/razshare/frizzante"
	"main/lib/config"
)

var guards = []f.Guard{
	config.GuardExpired,
	config.GuardVerified,
}

func init() {
	config.Server.LoadController(func(controller *f.Controller) {
		controller.
			WithBase(guards, base)
	})
}

type data struct {
	AccountId   string `json:"accountId"`
	DisplayName string `json:"displayName"`
}

func base(req *f.Request, res *f.Response) {
	session := f.SessionStart(req, res, config.SessionAdapter)

	fetchAccount, closeFetch := config.FindAccountById(session.Data.AccountId)
	defer closeFetch()

	var accountId string
	var displayName string
	var createdAt int
	var updatedAt int

	fetchAccount(&accountId, &displayName, &createdAt, &updatedAt)

	res.SendView(f.NewViewWithData(f.RenderModeFull, data{
		AccountId:   accountId,
		DisplayName: displayName,
	}))
	session.Save()
}
