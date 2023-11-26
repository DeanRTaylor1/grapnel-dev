package handlers

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/DeanRTaylor1/deans-site/logger"
)

//go:embed templates/*.html templates/common/*.html
var content embed.FS

func ServeHome(w http.ResponseWriter, r *http.Request, logger *logger.Logger) {
	logger.Debug("Accessed route: '/'")

	tmpl, err := template.ParseFS(content, "templates/*.html", "templates/common/*.html")
	if err != nil {
		logger.Error(fmt.Sprintf("Error rendering HTML template: %s", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	data := PageData{
		Title: "Sys.D Solutions - Home",
	}

	w.Header().Set("Content-Type", ContentTypeHTML)

	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		logger.Error(fmt.Sprintf("Error rendering HTML template: %s", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
