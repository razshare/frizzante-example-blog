package sessions

import (
	"encoding/json"
	"github.com/razshare/frizzante"
	"main/lib"
)

func Adapter(session *frizzante.Session[SessionData]) {
	session.WithExistsHandler(func() bool {
		return Archive.Has(session.Id, Key)
	})

	session.WithLoadHandler(func() {
		data := Archive.Get(session.Id, Key)
		unmarshalError := json.Unmarshal(data, &session.Data)
		if nil != unmarshalError {
			lib.Notifier.SendError(unmarshalError)
		}
	})

	session.WithSaveHandler(func() {
		data, marshalError := json.Marshal(session.Data)
		if nil != marshalError {
			lib.Notifier.SendError(marshalError)
			return
		}
		Archive.Set(session.Id, Key, data)
	})

	session.WithDestroyHandler(func() {
		Archive.RemoveDomain(session.Id)
	})

	if session.Exists() {
		session.Load()
		return
	}

	session.Data = NewSessionData()
}
