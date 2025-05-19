package pages

import f "github.com/razshare/frizzante"

type DefaultController struct {
	f.PageController
}

func (_ DefaultController) Configure() f.PageConfiguration {
	return f.PageConfiguration{
		Path:         "/",
		TryFileFirst: true,
	}
}

func (_ DefaultController) Base(request *f.Request, response *f.Response) {
	response.SendView(f.NewView(f.RenderModeFull))
}

func (_ DefaultController) Action(request *f.Request, response *f.Response) {
	response.SendView(f.NewView(f.RenderModeFull))
}
