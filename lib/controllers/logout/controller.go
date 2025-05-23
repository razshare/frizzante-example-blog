package logout

import f "github.com/razshare/frizzante"

import "main/lib/config"

var guards []f.Guard

func init() {
	config.Server.LoadController(func(controller *f.Controller) {
		controller.
			WithBase(guards, base).
			WithAction(guards, action)
	})
}

func base(req *f.Request, res *f.Response) {
	session := f.SessionStart(req, res, config.SessionAdapter)
	session.Data.Verified = false
	session.Save()
	res.SendNavigate("login")
}

func action(req *f.Request, res *f.Response) {
	session := f.SessionStart(req, res, config.SessionAdapter)
	session.Data.Verified = false
	session.Save()
	res.SendNavigate("login")
}
