package pages

import (
	"crypto/sha256"
	"fmt"
	frz "github.com/razshare/frizzante"
	"main/schemas"
	"time"
)

func Register(_ *frz.Server, req *frz.Request, _ *frz.Response, p *frz.Page) {
	frz.PageWithRenderMode(p, frz.RenderModeServer)

	form := frz.ReceiveForm(req)

	accountId := form.Get("accountId")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(form.Get("password"))))

	// Do nothing if it's not a submission.
	if "" == accountId {
		return
	}

	// Try login.
	now := time.Now().Unix()
	frz.SqlExecute(schemas.Sql, "insert into Account values(?,?,?,?)", accountId, password, now, now)
}
