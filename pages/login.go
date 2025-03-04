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

	logged := get("logged", false).(bool)

	// User is already logged in.
	if logged {
		frz.PageWithData(p, "message", "You're already logged in.")
		return
	}

	form := frz.ReceiveForm(req)

	submit := form.Has("submit")

	// Do nothing if it's not a submission.
	if !submit {
		return
	}

	id := form.Get("id")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

	// Try login.
	fetch, _ := frz.SqlFind(database.Sql, "select id from `User` where id = ? and password = ? limit 1", id, password)
	if !fetch(&id) {
		frz.PageWithData(p, "error", "Invalid credentials")
		return
	}

	frz.PageWithData(p, "message", "Login successful!")

	set("logged", true)
}
