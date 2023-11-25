package handlers

import (
	"embed"
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/DeanRTaylor1/deans-site/logger"
)

//go:embed images/*jpg
var imageFiles embed.FS

func ServeImages(w http.ResponseWriter, r *http.Request, logger *logger.Logger) {
	imageFilename := path.Base(r.URL.Path)
	fmt.Println(imageFilename)

	imageFile, err := imageFiles.ReadFile("images/" + imageFilename)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		logger.Error(fmt.Sprintf("Error reading embedded image file: %s", err.Error()))
		return
	}

	cacheDuration := 24 * time.Hour 
	SetCacheHeaders(w, ContentTypeJPG, cacheDuration, imageFilename)


	w.WriteHeader(http.StatusOK)
	w.Write(imageFile)
}
