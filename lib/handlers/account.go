package handlers

import (
	"context"
	frz "github.com/razshare/frizzante"
	"main/lib"
	"main/lib/database"
)

func GetAccount(c *frz.Connection) {
	state, _ := frz.Session(c, lib.State{})
	account, accountError := database.Queries.SqlFindAccountById(context.Background(), state.AccountId)
	if nil != accountError {
		c.SendView(frz.View{Name: "Account", Error: accountError.Error()})
		return
	}
	c.SendView(frz.View{Name: "Account", Data: account})
}
