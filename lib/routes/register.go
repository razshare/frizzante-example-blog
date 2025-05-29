package routes

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/razshare/frizzante"
	"main/lib/database"
	"main/lib/generated"
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

	_, accountError := database.Queries.SqlFindAccountById(context.Background(), id)
	if nil == accountError {
		res.SendView(frizzante.View{Name: "Register", Error: fmt.Errorf("account `%s` already exists", id)})
		return
	}

	addError := database.Queries.SqlAddAccount(context.Background(), generated.SqlAddAccountParams{
		ID:          id,
		DisplayName: displayName,
		Password:    password,
	})

	if nil != addError {
		res.SendView(frizzante.View{Name: "Register", Error: addError})
		return
	}

	res.SendNavigate("/login")
}
