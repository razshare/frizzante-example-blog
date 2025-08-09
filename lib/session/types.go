package session

import "time"

type Session struct {
	LastActivity time.Time
	Verified     bool
	Expired      bool
	AccountId    string
}
