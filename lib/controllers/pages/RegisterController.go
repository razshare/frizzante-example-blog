package pages

import (
	"crypto/sha256"
	"fmt"
	f "github.com/razshare/frizzante"
	"main/lib"
)

type RegisterData struct {
	Error string `json:"error"`
}

type RegisterController struct {
	f.PageController
}

func (_ RegisterController) Configure() f.PageConfiguration {
	return f.PageConfiguration{
		Path: "/register",
	}
}

func (_ RegisterController) Base(request *f.Request, response *f.Response) {
	response.SendView(f.NewView(false))
}

func (_ RegisterController) Action(request *f.Request, response *f.Response) {
	form := request.ReceiveForm()
	id := form.Get("id")
	if lib.AccountExists(id) {
		response.SendView(f.NewView(RegisterData{
			Error: fmt.Sprintf("Account %s already exists.", id),
		}))
		return
	}

	displayName := form.Get("displayName")
	rawPassword := form.Get("password")

	if "" == id || "" == displayName || "" == rawPassword {
		response.SendView(f.NewView(RegisterData{
			Error: "Please fill all fields.",
		}))
		return
	}

	password := fmt.Sprintf("%x", sha256.Sum256([]byte(rawPassword)))
	lib.AddAccount(id, displayName, password)
	response.SendNavigate("Login", f.NewView(false))
}
