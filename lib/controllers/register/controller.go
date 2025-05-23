package register

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
			WithBase(guards, base).
			WithAction(guards, action)
	})
}

type data struct {
	Error string `json:"error"`
}

func base(req *f.Request, res *f.Response) {
	res.SendView(f.NewView(f.RenderModeFull))
}

func action(req *f.Request, res *f.Response) {
	form := req.ReceiveForm()
	id := form.Get("id")
	if config.AccountExists(id) {
		res.SendView(f.NewViewWithData(f.RenderModeFull, data{
			Error: fmt.Sprintf("Account %s already exists.", id),
		}))
		return
	}

	displayName := form.Get("displayName")
	rawPassword := form.Get("password")

	if "" == id || "" == displayName || "" == rawPassword {
		res.SendView(f.NewViewWithData(f.RenderModeFull, data{
			Error: "Please fill all fields.",
		}))
		return
	}

	password := fmt.Sprintf("%x", sha256.Sum256([]byte(rawPassword)))
	config.AddAccount(id, displayName, password)
	res.SendNavigate("login")
}
