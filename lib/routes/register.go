package routes

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/razshare/frizzante"
	"main/lib/database"
	"main/lib/generated"
	"main/lib/value"
)

func GetRegister(req *frizzante.Request, res *frizzante.Response) {
	res.SendView(frizzante.View{Name: "Register"})
}

func PostRegister(req *frizzante.Request, res *frizzante.Response) {
	form := req.ReceiveForm()
	id := form.Get("id")

	displayName := form.Get("displayName")
	rawPassword := form.Get("password")

	if "" == id || "" == displayName || "" == rawPassword {
		res.SendView(frizzante.View{Name: "Register", Error: errors.New("please fill all fields")})
		return
	}

	password := fmt.Sprintf("%x", sha256.Sum256([]byte(rawPassword)))

	account := value.Wrap(database.Queries.SqlFindAccountById(context.Background(), id))
	if !account.Ok() {
		res.SendView(frizzante.View{Name: "Register", Error: account.Error})
		return
	}

	addAttempt := value.WrapNothing(database.Queries.SqlAddAccount(context.Background(), generated.SqlAddAccountParams{
		ID:          id,
		DisplayName: displayName,
		Password:    password,
	}))

	if !addAttempt.Ok() {
		res.SendView(frizzante.View{Name: "Register", Error: addAttempt.Error})
		return
	}

	res.SendNavigate("/login")
}
