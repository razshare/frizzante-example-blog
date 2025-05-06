package pages

import (
	"crypto/sha256"
	"fmt"
	f "github.com/razshare/frizzante"
	"main/lib/sql"
)

func registerActionFunction(request *f.Request, response *f.Response, _ *f.View) {
	form := f.ReceiveForm(request)
	id := form.Get("id")
	displayName := form.Get("displayName")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))
	sql.AddAccount(id, displayName, password)
	f.SendNavigate(response, "Login")
}

func Register(context f.PageContext) {
	path, view, _, action := context()
	path("/register")
	view(f.ViewReference("Register"))
	action(registerActionFunction)
}
