package routes

import (
	"context"
	"github.com/razshare/frizzante"
	"main/lib/database"
	"main/lib/guards"
	"main/lib/sessions"
)

func GetAccount(req *frizzante.Request, res *frizzante.Response) {
	if !frizzante.AllGuardsPass(req, res, guards.NotExpired, guards.Verified) {
		return
	}

	session := frizzante.SessionStart(req, res, sessions.Adapter)
	account, accountError := database.Queries.SqlFindAccountById(context.Background(), session.Data.AccountId)

	if nil != accountError {
		res.SendView(frizzante.View{Name: "Account", Error: accountError})
		return
	}

	res.SendView(frizzante.View{Name: "Account", Data: account})
}
