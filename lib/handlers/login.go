package handlers

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	frz "github.com/razshare/frizzante"
	"main/lib"
	"main/lib/database"
	"main/lib/generated"
	"time"
)

func GetLogin(c *frz.Connection) {
	c.SendView(frz.View{Name: "Login"})
}

func PostLogin(c *frz.Connection) {
	state, operator := frz.Session(c, lib.State{})
	defer operator.Save(state)
	form := c.ReceiveForm()
	id := form.Get("id")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

	_, accountError := database.Queries.SqlVerifyAccount(context.Background(), generated.SqlVerifyAccountParams{
		ID:       id,
		Password: password,
	})

	if nil != accountError {
		c.SendView(frz.View{Name: "Login", Error: errors.New("invalid credentials")})
		return
	}

	state.LastActivity = time.Now()
	state.Verified = true
	state.AccountId = id

	c.SendNavigate("/board")
}
