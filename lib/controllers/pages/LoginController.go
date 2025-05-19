package pages

import (
	"crypto/sha256"
	"fmt"
	f "github.com/razshare/frizzante"
	"main/lib"
	"main/lib/sessions"
)

type LoginData struct {
	Error string `json:"error"`
}

type LoginController struct {
	f.PageController
}

func (_ LoginController) Configure() f.PageConfiguration {
	return f.PageConfiguration{
		Path: "/login",
	}
}

func (_ LoginController) Base(_ *f.Request, response *f.Response) {
	response.SendView(f.NewView(f.RenderModeFull))
}

func (_ LoginController) Action(request *f.Request, response *f.Response) {
	form := request.ReceiveForm()
	id := form.Get("id")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

	if !lib.VerifyAccount(id, password) {
		response.SendView(f.NewViewWithData(f.RenderModeFull, LoginData{
			Error: "Invalid credentials",
		}))
		return
	}

	session := f.SessionStart(request, response, sessions.Archived)
	session.Data.Verified = true
	session.Data.AccountId = id
	session.Save()
	response.SendNavigate("Board")
}
