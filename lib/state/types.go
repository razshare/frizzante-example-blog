package state

import "time"

type State struct {
	LastActivity time.Time
	Verified     bool
	Expired      bool
	AccountId    string
}
