package pages

import (
	"crypto/sha256"
	"fmt"
	f "github.com/razshare/frizzante"
	"main/lib"
)

var Register = f.NewPageController().WithBase(registerBase).WithAction(registerAction)

type RegisterData struct {
	Error string `json:"error"`
}

func registerBase(req *f.Request, res *f.Response) {
	res.SendView(f.NewView(f.RenderModeFull))
}

func registerAction(req *f.Request, res *f.Response) {
	form := req.ReceiveForm()
	id := form.Get("id")
	if lib.AccountExists(id) {
		res.SendView(f.NewViewWithData(f.RenderModeFull, RegisterData{
			Error: fmt.Sprintf("Account %s already exists.", id),
		}))
		return
	}

	displayName := form.Get("displayName")
	rawPassword := form.Get("password")

	if "" == id || "" == displayName || "" == rawPassword {
		res.SendView(f.NewViewWithData(f.RenderModeFull, RegisterData{
			Error: "Please fill all fields.",
		}))
		return
	}

	password := fmt.Sprintf("%x", sha256.Sum256([]byte(rawPassword)))
	lib.AddAccount(id, displayName, password)
	res.SendNavigate("Login")
}
