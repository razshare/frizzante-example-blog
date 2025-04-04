package pages

import (
	"crypto/sha256"
	"fmt"
	frz "github.com/razshare/frizzante"
	"main/lib"
)

func Login(_ *frz.Server, req *frz.Request, res *frz.Response, p *frz.Page) {
	frz.PageWithRenderMode(p, frz.RenderModeServer)

	get, set, _ := frz.SessionStart(req, res)
	form := frz.ReceiveForm(req)

	logged := get("Logged", false).(bool)

	// User is already logged in.
	if logged {
		// Logout...
		logout := form.Has("Logout")
		if logout {
			set("Logged", false)
			return
		}

		// ...or hint to the user to logout.
		frz.PageWithData(p, "Message", "You're already logged in.")
		frz.PageWithData(p, "Logged", true)
		return
	}

	id := form.Get("Id")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("Password"))))

	// Do nothing if it's not a submission.
	if "" == id {
		return
	}

	// Verify login info.
	if !lib.VerifyAccount(id, password) {
		frz.PageWithData(p, "Error", "Invalid credentials")
		return
	}

	frz.PageWithData(p, "Message", "Login successful!")
	frz.PageWithData(p, "Logged", true)

	set("Logged", true)
}
