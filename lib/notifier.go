package lib

import f "github.com/razshare/frizzante"

var Notifier = f.NotifierCreate()

// Failure sends and error to the notifier.
func Failure(err error) {
	f.NotifierSendError(Notifier, err)
}

// Info sends a message to the notifier.
func Info(message string) {
	f.NotifierSendMessage(Notifier, message)
}
