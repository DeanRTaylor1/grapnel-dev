package handlers

import (
	"embed"
	"fmt"
	"net/http"
	"path"

	"github.com/DeanRTaylor1/deans-site/constants"
	"github.com/DeanRTaylor1/deans-site/logger"
)

//go:embed images/*jpg
var imageFiles embed.FS

func ServeImages(w http.ResponseWriter, r *http.Request, logger *logger.Logger) {
	imageFilename := path.Base(r.URL.Path)

	imageFile, err := imageFiles.ReadFile("images/" + imageFilename)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		logger.Error(fmt.Sprintf("Error reading embedded image file: %s", err.Error()))
		return
	}

	SetCacheHeaders(w, ContentTypeJPG, constants.CacheDuration, imageFilename)

	w.WriteHeader(http.StatusOK)
	w.Write(imageFile)
}

//go:embed images/icons/*png  images/icons/*ico
var iconFiles embed.FS

func ServeIcons(w http.ResponseWriter, r *http.Request, logger *logger.Logger) {
	iconFilename := path.Base(r.URL.Path)

	iconFile, err := iconFiles.ReadFile("images/icons/" + iconFilename)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		logger.Error(fmt.Sprintf("Error reading embedded icon file: %s", err.Error()))
		return
	}

	SetCacheHeaders(w, "image/png", constants.CacheDuration, iconFilename)

	w.Header().Set("Content-Type", ContentTypePNG)
	w.WriteHeader(http.StatusOK)
	w.Write(iconFile)
}

func ServeFavicon(w http.ResponseWriter, r *http.Request, logger *logger.Logger) {
	iconFilename := path.Base(r.URL.Path)

	iconFile, err := iconFiles.ReadFile("images/icons/" + iconFilename)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		logger.Error(fmt.Sprintf("Error reading embedded icon file: %s", err.Error()))
		return
	}

	SetCacheHeaders(w, ContentTypeICO, constants.CacheDuration, iconFilename)

	w.WriteHeader(http.StatusOK)
	w.Write(iconFile)
}
