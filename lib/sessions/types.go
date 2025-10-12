package sessions

import "time"

type Session struct {
	LastActivity  time.Time
	LoggedIn      bool
	LoginExpired  bool
	AccountId     string
	RegisterError string
	LoginError    string
	FormError     string
	BoardError    string
}
