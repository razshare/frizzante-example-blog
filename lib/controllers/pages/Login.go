package pages

import (
	"crypto/sha256"
	"fmt"
	f "github.com/razshare/frizzante"
	"main/lib"
	"main/lib/sessions"
)

var Login = f.NewPageController().WithAction(loginAction)

type LoginData struct {
	Error string `json:"error"`
}

func loginAction(req *f.Request, res *f.Response) {
	form := req.ReceiveForm()
	id := form.Get("id")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

	if !lib.VerifyAccount(id, password) {
		res.SendView(f.NewViewWithData(f.RenderModeFull, LoginData{
			Error: "Invalid credentials",
		}))
		return
	}

	session := f.SessionStart(req, res, sessions.Archived)
	session.Data.Verified = true
	session.Data.AccountId = id
	session.Save()
	res.SendNavigate("Board")
}
