package render

import (
	"embed"
	"log"

	"main/lib/core/views"
)

type Config struct {
	App      string
	Efs      embed.FS
	Limit    int
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

type Render = func(view views.View) (html string, err error)
