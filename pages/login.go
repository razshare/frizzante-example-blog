package pages

import (
	"crypto/sha256"
	"fmt"
	frz "github.com/razshare/frizzante"
	"main/database"
)

func Login(_ *frz.Server, req *frz.Request, res *frz.Response, p *frz.Page) {
	frz.PageWithRenderMode(p, frz.RenderModeServer)

	get, set, _ := frz.SessionStart(req, res)
	form := frz.ReceiveForm(req)

	logged := get("logged", false).(bool)

	// User is already logged in.
	if logged {
		// Logout...
		logout := form.Has("logout")
		if logout {
			set("logged", false)
			return
		}

		// ...or hint to the user to logout.
		frz.PageWithData(p, "message", "You're already logged in.")
		frz.PageWithData(p, "logged", true)
		return
	}

	id := form.Get("id")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

	// Do nothing if it's not a submission.
	if "" == id {
		return
	}

	// Try login.
	fetch, _ := frz.SqlFind(database.Sql, "select Id from Account where Id = ? and Password = ? limit 1", id, password)
	if !fetch(&id) {
		frz.PageWithData(p, "error", "Invalid credentials")
		return
	}

	frz.PageWithData(p, "message", "Login successful!")
	frz.PageWithData(p, "logged", true)

	set("logged", true)
}
