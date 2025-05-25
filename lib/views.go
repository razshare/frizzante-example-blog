package lib

import "github.com/razshare/frizzante"

func NewErrorView(err error) frizzante.View {
	return frizzante.View{
		Name: "Error",
		Data: map[string]string{
			"error": err.Error(),
		},
	}
}
