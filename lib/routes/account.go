package routes

import (
	"context"
	"github.com/razshare/frizzante"
	"main/lib"
	"main/lib/database"
	"main/lib/guards"
	"main/lib/sessions"
	"main/lib/value"
)

func init() {
	lib.Server.WithRequestHandler("GET /account", GetAccount)
}

func GetAccount(req *frizzante.Request, res *frizzante.Response) {
	if !frizzante.AllGuardsPass(req, res, guards.NotExpired, guards.Verified) {
		return
	}

	session := frizzante.SessionStart(req, res, sessions.Adapter)
	account := value.Wrap(database.Queries.SqlFindAccountById(context.Background(), session.Data.AccountId))

	if !account.Ok() {
		res.SendView(frizzante.View{Name: "Account", Error: account.Error})
		return
	}

	res.SendView(frizzante.View{Name: "Account", Data: account.Value})
}
