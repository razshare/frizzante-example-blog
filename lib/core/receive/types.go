package receive

import (
	"mime/multipart"
	"net/url"

	"main/lib/core/client"
)

type MultipartFormFile struct {
	multipart.File
	multipart.FileHeader
}

type MultipartForm struct {
	url.Values
	Client *client.Client
}
