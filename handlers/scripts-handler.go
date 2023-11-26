package handlers

import (
	"embed"
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/DeanRTaylor1/deans-site/config"
	"github.com/DeanRTaylor1/deans-site/logger"
)

//go:embed scripts/*.js
var scripts embed.FS

func ServeScripts(w http.ResponseWriter, r *http.Request, logger *logger.Logger, config config.EnvConfig) {
	scriptName := path.Base(r.URL.Path)

	scriptFile, err := scripts.ReadFile("scripts/" + scriptName)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		logger.Error(fmt.Sprintf("Error reading embedded font file: %s", err.Error()))
		return
	}

	if config.IsProduction {
		cacheDuration := 24 * time.Hour
		SetCacheHeaders(w, ContentTypeJavaScript, cacheDuration, scriptName)
	}

	w.Header().Set("Content/Type", ContentTypeJavaScript)
	w.WriteHeader(http.StatusOK)
	w.Write(scriptFile)
}
