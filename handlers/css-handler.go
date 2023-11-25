package handlers

import (
	"embed"
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/DeanRTaylor1/deans-site/logger"
)

//go:embed styles/output.css
var cssContent embed.FS

func ServeCss(w http.ResponseWriter, r *http.Request, logger *logger.Logger) {
	logger.Debug("Accessed route: '/output.css'")

	filePath := "styles/output.css"
	styles, err := cssContent.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		logger.Error(fmt.Sprintf("Error reading embedded CSS file: %s", err.Error()))
		return
	}	
	
	cacheDuration := 24 * time.Hour 
	SetCacheHeaders(w, ContentTypeCSS, cacheDuration, "output.css")

	w.WriteHeader(http.StatusOK)
	w.Write(styles)
}

//go:embed templates/fonts/*.otf
var fontFiles embed.FS

func ServeFonts(w http.ResponseWriter, r *http.Request, logger *logger.Logger) {
	fontFilename := path.Base(r.URL.Path)

	fontFile, err := fontFiles.ReadFile("templates/fonts/" + fontFilename)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		logger.Error(fmt.Sprintf("Error reading embedded font file: %s", err.Error()))
		return
	}

	cacheDuration := 24 * time.Hour 
	SetCacheHeaders(w, ContentTypeFontOpen, cacheDuration, fontFilename)


	w.WriteHeader(http.StatusOK)
	w.Write(fontFile)
}
