package config

import (
	"encoding/json"
	f "github.com/razshare/frizzante"
	"time"
)

type SessionData struct {
	Items        []Item    `json:"items"`
	LastActivity time.Time `json:"lastActivity"`
	Verified     bool      `json:"verified"`
	Expired      bool      `json:"expired"`
	AccountId    string    `json:"accountId"`
}

type Item struct {
	Checked     bool   `json:"checked"`
	Description string `json:"description"`
}

func NewSessionData() SessionData {
	return SessionData{
		Items: []Item{
			{Checked: false, Description: "Pet the cat."},
			{Checked: false, Description: "Do laundry"},
			{Checked: false, Description: "Pet the cat."},
			{Checked: false, Description: "Cook"},
			{Checked: false, Description: "Pet the cat."},
		},
		LastActivity: time.Now(),
		Expired:      false,
		Verified:     true,
	}
}

var key = "session.json"
var notifier = f.NewNotifier()
var archive = f.NewArchiveOnDisk(".sessions", time.Second/2)

func SessionAdapter(session *f.Session[SessionData]) {
	session.WithExistsHandler(func() bool {
		return archive.Has(session.Id, key)
	})

	session.WithLoadHandler(func() {
		data := archive.Get(session.Id, key)
		unmarshalError := json.Unmarshal(data, &session.Data)
		if nil != unmarshalError {
			notifier.SendError(unmarshalError)
		}
	})

	session.WithSaveHandler(func() {
		data, marshalError := json.Marshal(session.Data)
		if nil != marshalError {
			notifier.SendError(marshalError)
			return
		}
		archive.Set(session.Id, key, data)
	})

	session.WithDestroyHandler(func() {
		archive.RemoveDomain(session.Id)
	})

	if session.Exists() {
		session.Load()
		return
	}

	session.Data = NewSessionData()
}
