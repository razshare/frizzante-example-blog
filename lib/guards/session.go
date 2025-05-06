package guards

import f "github.com/razshare/frizzante"

func Session(context f.GuardContext) {
	handler := context()
	handler(func(request *f.Request, response *f.Response, pass func()) {
		f.SessionStart(request, response)
		pass()
	})
}
