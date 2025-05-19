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

func (_ DefaultController) Base(_ *f.Request, response *f.Response) {
	response.SendView(f.NewView(false))
}

func (_ DefaultController) Action(_ *f.Request, response *f.Response) {
	response.SendView(f.NewView(false))
}
