package server

import (
	"embed"
	"log"
	"net/http"

	"main/lib/core/guard"
	"main/lib/core/route"
	view "main/lib/core/view"
)

type Server struct {
	*http.Server
	Guards      []guard.Guard
	Routes      []route.Route
	PublicRoot  string
	SecureAddr  string
	Certificate string
	Key         string
	Channels    Channels
	InfoLog     *log.Logger
	Efs         embed.FS
	Cors        *http.CrossOriginProtection
	Render      func(view view.View) (html string, err error)
}

type Channels struct {
	Stop chan any
}
