package pages

import (
	"crypto/sha256"
	"fmt"
	frz "github.com/razshare/frizzante"
	"main/database"
	"time"
)

func Register(_ *frz.Server, req *frz.Request, res *frz.Response, p *frz.Page) {
	form := frz.ReceiveForm(req)

	submit := form.Has("submit")

	// Do nothing if it's not a submission.
	if !submit {
		return
	}

	id := form.Get("id")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

	// Try login.
	now := time.Now().Unix()
	frz.SqlExecute(database.Sql, "insert into `User` values(?,?,?,?)", id, password, now, now)
}
