package pages

import f "github.com/razshare/frizzante"

type RedirectData struct {
}

type RedirectController struct {
	f.PageController
}

func (_ RedirectController) Configure() f.PageConfiguration {
	return f.PageConfiguration{
		Path: "/redirect",
	}
}

func (_ RedirectController) Base(request *f.Request, response *f.Response) {
	response.SendNavigateWithQuery("Expired", "?asd=1")
}

func (_ RedirectController) Action(request *f.Request, response *f.Response) {
	response.SendNavigateWithQuery("Expired", "?asd=1")
}
