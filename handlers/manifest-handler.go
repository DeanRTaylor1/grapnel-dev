package handlers

import (
	"embed"
	"fmt"
	"net/http"
	"path"

	"github.com/DeanRTaylor1/deans-site/constants"
	"github.com/DeanRTaylor1/deans-site/logger"
)

//go:embed manifest/manifest.json
var manifestFiles embed.FS

func ServeManifest(w http.ResponseWriter, r *http.Request, logger *logger.Logger) {
	manifestFile := path.Base(r.URL.Path)

	imageFile, err := manifestFiles.ReadFile("manifest/" + manifestFile)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		logger.Error(fmt.Sprintf("Error reading embedded image file: %s", err.Error()))
		return
	}

	SetCacheHeaders(w, ContentTypeJSON, constants.CacheDuration, manifestFile)

	w.WriteHeader(http.StatusOK)
	w.Write(imageFile)
}
