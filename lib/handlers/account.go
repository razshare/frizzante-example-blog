package handlers

import (
	"context"
	"github.com/razshare/frizzante"
	"main/lib"
	"main/lib/database"
)

func GetAccount(c *frizzante.Connection) {
	lib.SessionStartProtected(c, func(state *lib.State) {
		account, accountError := database.Queries.SqlFindAccountById(context.Background(), state.AccountId)
		if nil != accountError {
			c.SendView(frizzante.View{Name: "Account", Error: accountError})
			return
		}
		c.SendView(frizzante.View{Name: "Account", Data: account})
	})
}
