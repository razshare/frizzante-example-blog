package receive

import (
	"mime/multipart"
	"net/url"

	"main/lib/core/clients"
)

type MultipartFormFile struct {
	multipart.File
	multipart.FileHeader
}

type MultipartForm struct {
	url.Values
	Client *clients.Client
}
