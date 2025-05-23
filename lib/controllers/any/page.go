package any

import (
	f "github.com/razshare/frizzante"
	"main/lib/config"
)

var guards = []f.Guard{
	config.GuardExpired,
	config.GuardVerified,
}

func init() {
	config.Server.LoadController(func(controller *f.Controller) {
		controller.
			WithBase(guards, base).
			WithAction(guards, action)
	})
}

func base(req *f.Request, res *f.Response) {
	res.SendView(f.NewView(f.RenderModeFull))
}

func action(req *f.Request, res *f.Response) {
	res.SendView(f.NewView(f.RenderModeFull))
}
