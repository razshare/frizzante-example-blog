package pages

import (
	"crypto/sha256"
	"fmt"
	frz "github.com/razshare/frizzante"
	"main/database"
	"time"
)

func Register(_ *frz.Server, req *frz.Request, _ *frz.Response, p *frz.Page) {
	frz.PageWithRenderMode(p, frz.RenderModeServer)

	form := frz.ReceiveForm(req)

	id := form.Get("id")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

	// Do nothing if it's not a submission.
	if "" == id {
		return
	}

	// Try login.
	now := time.Now().Unix()
	frz.SqlExecute(database.Sql, "insert into `User` values(?,?,?,?)", id, password, now, now)
}
