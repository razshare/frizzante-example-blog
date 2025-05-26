package sessions

import "time"

type SessionData struct {
	LastActivity time.Time `json:"lastActivity"`
	Verified     bool      `json:"verified"`
	Expired      bool      `json:"expired"`
	AccountId    string    `json:"accountId"`
}

func NewSessionData() SessionData {
	return SessionData{
		LastActivity: time.Now(),
		Expired:      false,
		Verified:     true,
	}
}
