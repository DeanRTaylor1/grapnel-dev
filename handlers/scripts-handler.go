package handlers

import (
	"embed"
	"fmt"
	"net/http"
	"path"

	"github.com/DeanRTaylor1/deans-site/config"
	"github.com/DeanRTaylor1/deans-site/constants"
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

	SetCacheHeaders(w, ContentTypeJavaScript, constants.CacheDuration, scriptName)

	w.Write(scriptFile)
}
