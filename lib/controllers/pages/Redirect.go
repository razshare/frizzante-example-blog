package pages

import f "github.com/razshare/frizzante"

var Redirect = f.NewPageController().WithBase(redirectBase).WithAction(redirectAction)

func redirectBase(req *f.Request, res *f.Response) {
	res.SendNavigateWithQuery("Expired", "?asd=1")
}

func redirectAction(req *f.Request, res *f.Response) {
	res.SendNavigateWithQuery("Expired", "?asd=1")
}
