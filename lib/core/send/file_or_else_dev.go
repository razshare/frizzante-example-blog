//go:build dev

package send

import (
	"net/http"
	"path/filepath"
	"strings"

	"main/lib/core/clients"
	"main/lib/core/files"
	"main/lib/core/mime"
	"main/lib/core/stacks"
)

// FileOrElse sends the file requested by the client, or else falls back.
func FileOrElse(client *clients.Client, orElse func()) {
	if client.WebSocket != nil {
		client.Config.ErrorLog.Println("file_or_else does not support web sockets", stacks.Trace())
		return
	}

	if client.EventName != "" {
		client.Config.ErrorLog.Println("file_or_else does not support server sent events", stacks.Trace())
		return
	}

	var name string

	if strings.HasPrefix(client.Request.RequestURI, "/") {
		name = filepath.Join(client.Config.PublicRoot, client.Request.RequestURI[1:])
	} else {
		name = filepath.Join(client.Config.PublicRoot, client.Request.RequestURI)
	}

	if files.IsFile(name) {
		if "" == client.Writer.Header().Get("Content-Type") {
			Header(client, "Content-Type", mime.Parse(name))
		}

		http.ServeFile(client.Writer, client.Request, name)
		return
	}

	orElse()
}
