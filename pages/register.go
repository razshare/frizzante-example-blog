package pages

import (
	"crypto/sha256"
	"fmt"
	frz "github.com/razshare/frizzante"
	"main/schemas"
	"time"
)

func Register(_ *frz.Server, req *frz.Request, res *frz.Response, p *frz.Page) {
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
	frz.SqlExecute(schemas.Sql, "insert into Account(Id,Password,CreatedAt,UpdatedAt) values(?,?,?,?)", id, password, now, now)

	frz.SendRedirectToPage(res, "login", map[string]string{})
}
