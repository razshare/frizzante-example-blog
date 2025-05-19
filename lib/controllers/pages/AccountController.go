package pages

import f "github.com/razshare/frizzante"

type AccountController struct {
	f.PageController
}

func (_ AccountController) Configure() f.PageConfiguration {
	return f.PageConfiguration{
		Path: "/account",
	}
}

func (_ AccountController) Base(request *f.Request, response *f.Response) {
	response.SendView(f.NewView(f.RenderModeFull))
}

func (_ AccountController) Action(request *f.Request, response *f.Response) {
	response.SendView(f.NewView(f.RenderModeFull))
}
