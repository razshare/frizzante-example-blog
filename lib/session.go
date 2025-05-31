package lib

import "github.com/razshare/frizzante"

func SessionStartProtected(connection *frizzante.Connection, handler frizzante.SessionHandler[State]) {
	frizzante.NewSession[State](connection).WithGuards(protected).Start(handler)
}

func SessionStartPublic(connection *frizzante.Connection, handler frizzante.SessionHandler[State]) {
	frizzante.NewSession[State](connection).Start(handler)
}
