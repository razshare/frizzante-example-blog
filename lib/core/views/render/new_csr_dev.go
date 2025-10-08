//go:build dev && no_js_runtime

package render

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"main/lib/core/views"
)

func New(conf Config) Render {
	var app = conf.App

	if app == "" {
		app = "app"
	}

	var index = filepath.Join(app, "dist", "client", "index.html")

	index = strings.ReplaceAll(index, "/", string(filepath.Separator))
	index = strings.ReplaceAll(index, "\\", string(filepath.Separator))

	return func(view views.View) (document string, err error) {
		var indexData []byte
		if indexData, err = os.ReadFile(index); err != nil {
			return
		}

		document = string(indexData)

		var data []byte
		if data, err = json.Marshal(views.NewData(view)); err != nil {
			return "", err
		}

		document = strings.Replace(document, "<!--app-head-->", fmt.Sprintf(HeadFormat, view.Title), 1)
		document = strings.Replace(document, "<!--app-body-->", fmt.Sprintf(BodyFormat, ""), 1)
		document = strings.Replace(document, "<!--app-data-->", fmt.Sprintf(DataFormat, data), 1)

		return document, nil
	}
}
