package config

import f "github.com/razshare/frizzante"

var Notifier = f.NewNotifier()

// Failure sends and error to the notifier.
func Failure(err error) {
	Notifier.SendError(err)
}

// Info sends a message to the notifier.
func Info(message string) {
	Notifier.SendMessage(message)
}
