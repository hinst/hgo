package hgo

import (
	"os"
	"path/filepath"
)

var AppDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
