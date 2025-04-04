package pages

import (
	"crypto/sha256"
	"fmt"
	frz "github.com/razshare/frizzante"
	"main/lib"
)

func Register(_ *frz.Server, req *frz.Request, res *frz.Response, p *frz.Page) {
	frz.PageWithRenderMode(p, frz.RenderModeServer)

	form := frz.ReceiveForm(req)

	id := form.Get("Id")
	displayName := form.Get("DisplayName")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("Password"))))

	// Do nothing if it's not a submission.
	if "" == id {
		return
	}

	// Register.
	lib.AddAccount(id, displayName, password)

	frz.SendRedirectToPage(res, "Login", map[string]string{})
}
