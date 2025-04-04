package lib

import frz "github.com/razshare/frizzante"

var Notifier = frz.NotifierCreate()

// Failure sends and error to the notifier.
func Failure(err error) {
	frz.NotifierSendError(Notifier, err)
}

// Info sends a message to the notifier.
func Info(message string) {
	frz.NotifierSendMessage(Notifier, message)
}
