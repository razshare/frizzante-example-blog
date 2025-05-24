package lib

import (
	f "github.com/razshare/frizzante"
	"time"
)

var archive = f.NewArchiveOnDisk(".sessions", time.Second/2)
