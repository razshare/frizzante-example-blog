package guards

import f "github.com/razshare/frizzante"

func Verified(request *f.Request, response *f.Response, pass func()) {
	session := f.SessionStart(request, response)
	if !f.SessionHas(session, "verified") || !f.SessionGet[bool](session, "verified") {
		f.ResponseSendNavigate(response, "Login")
		return
	}

	pass()
}
