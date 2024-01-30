package data

import (
	"path/filepath"
	"time"

	"github.com/adrg/xdg"
)

var DataFilename = filepath.Join(xdg.DataHome, "itsrock", "daggervers.json")

const (
	endpointURL    = "https://daggerverse.dev/api/refs"
	maxAgeDuration = 5 * time.Hour
)
