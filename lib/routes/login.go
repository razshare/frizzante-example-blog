package routes

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/razshare/frizzante"
	"main/lib/database"
	"main/lib/generated"
	"main/lib/sessions"
	"main/lib/value"
)

func GetLogin(req *frizzante.Request, res *frizzante.Response) {
	res.SendView(frizzante.View{Name: "Login"})
}

func PostLogin(req *frizzante.Request, res *frizzante.Response) {
	form := req.ReceiveForm()
	id := form.Get("id")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

	account := value.Wrap(database.Queries.SqlVerifyAccount(context.Background(), generated.SqlVerifyAccountParams{
		ID:       id,
		Password: password,
	}))

	if !account.Ok() {
		res.SendView(frizzante.View{Name: "Login", Error: errors.New("invalid credentials")})
		return
	}

	session := frizzante.SessionStart(req, res, sessions.Adapter)
	session.Data.Verified = true
	session.Data.AccountId = id
	session.Save()

	res.SendNavigate("/board")
}
