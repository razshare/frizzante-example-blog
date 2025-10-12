//go:build dev

package send

import (
	"net/http"
	"path/filepath"
	"strings"

	"main/lib/core/clients"
	"main/lib/core/files"
	"main/lib/core/mime"
	"main/lib/core/stack"
)

// RequestedFile sends the file requested by the client.
//
// Returns false if connection is web sockets, server sent events
// or the file was not found.
func RequestedFile(client *clients.Client) bool {
	if client.WebSocket != nil {
		client.Config.ErrorLog.Println("send.RequestedFile() does not support web sockets", stack.Trace())
		return false
	}

	if client.EventName != "" {
		client.Config.ErrorLog.Println("send.RequestedFile() does not support server sent events", stack.Trace())
		return false
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
		return true
	}

	return false
}
