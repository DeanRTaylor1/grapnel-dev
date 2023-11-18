package handlers

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/DeanRTaylor1/deans-site/logger"
)

//go:embed templates/*.html
var content embed.FS

func ServeHome(w http.ResponseWriter, r *http.Request, logger *logger.Logger) {
	logger.Debug("Accessed route: '/'")

	tmpl, err := template.ParseFS(content, "templates/*.html")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)

	err = tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		logger.Error(fmt.Sprintf("Error rendering HTML template: %s", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
