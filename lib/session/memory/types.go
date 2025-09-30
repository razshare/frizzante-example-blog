package memory

import "time"

type State struct {
	LastActivity time.Time
	LoggedIn     bool
	LoginExpired bool
	AccountId    string
}
