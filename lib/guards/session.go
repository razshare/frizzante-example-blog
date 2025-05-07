package guards

import f "github.com/razshare/frizzante"

func Session(request *f.Request, response *f.Response, pass func()) {
	session := f.SessionStart(request, response)
	verified :=
		f.SessionHas(session, "verified") &&
			f.SessionGet[bool](session, "verified")

	if !verified {
		f.ResponseSendNavigate(response, "Login")
	}

	pass()
}
