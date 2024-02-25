package editor

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed assets
var assetFS embed.FS

var Assets = func() *RuntimeFiles {
	fs, err := fs.Sub(assetFS, "assets")
	if err != nil {
		return nil
	}
	return NewRuntimeFiles(http.FS(fs))
}()
