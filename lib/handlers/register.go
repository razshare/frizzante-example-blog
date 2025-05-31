package handlers

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/razshare/frizzante"
	"main/lib/database"
	"main/lib/generated"
)

func GetRegister(c *frizzante.Connection) {
	c.SendView(frizzante.View{Name: "Register"})
}

func PostRegister(c *frizzante.Connection) {
	form := c.ReceiveForm()
	id := form.Get("id")

	displayName := form.Get("displayName")
	rawPassword := form.Get("password")

	if "" == id || "" == displayName || "" == rawPassword {
		c.SendView(frizzante.View{Name: "Register", Error: errors.New("please fill all fields")})
		return
	}

	password := fmt.Sprintf("%x", sha256.Sum256([]byte(rawPassword)))

	_, accountError := database.Queries.SqlFindAccountById(context.Background(), id)
	if nil == accountError {
		c.SendView(frizzante.View{Name: "Register", Error: fmt.Errorf("account `%s` already exists", id)})
		return
	}

	addError := database.Queries.SqlAddAccount(context.Background(), generated.SqlAddAccountParams{
		ID:          id,
		DisplayName: displayName,
		Password:    password,
	})

	if nil != addError {
		c.SendView(frizzante.View{Name: "Register", Error: addError})
		return
	}

	c.SendNavigate("/login")
}
