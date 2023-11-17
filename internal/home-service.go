package services

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/DeanRTaylor1/deans-site/logger"
)

func Home(w http.ResponseWriter, r *http.Request, logger *logger.Logger) {

	logger.Debug("Accessed route: '/'")

	tmplDir, err := os.Getwd()
	if err != nil {
		logger.Error(fmt.Sprintf("Error getting root dir: %s", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles(tmplDir + "/templates/index.html")
	if err != nil {
		logger.Error(fmt.Sprintf("Error parsing HTML template: %s", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = tmpl.Execute(w, nil)
	if err != nil {
		logger.Error(fmt.Sprintf("Error rendering HTML template: %s", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
