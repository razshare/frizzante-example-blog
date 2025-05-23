package login

import (
	"crypto/sha256"
	"fmt"
	f "github.com/razshare/frizzante"
	"main/lib/config"
)

var guards []f.Guard

func init() {
	config.Server.LoadController(func(controller *f.Controller) {
		controller.
			WithAction(guards, action)
	})
}

type data struct {
	Error string `json:"error"`
}

func action(req *f.Request, res *f.Response) {
	form := req.ReceiveForm()
	id := form.Get("id")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

	if !config.VerifyAccount(id, password) {
		res.SendView(f.NewViewWithData(f.RenderModeFull, data{
			Error: "Invalid credentials",
		}))
		return
	}

	session := f.SessionStart(req, res, config.SessionAdapter)
	session.Data.Verified = true
	session.Data.AccountId = id
	session.Save()
	res.SendNavigate("board")
}
