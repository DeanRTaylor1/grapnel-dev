package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/DeanRTaylor1/deans-site/constants"
	"github.com/DeanRTaylor1/deans-site/logger"
)

func ServeFaq(w http.ResponseWriter, r *http.Request, logger *logger.Logger) {
	logger.Debug("Accessed route: '/faq'")

	tmpl, err := template.ParseFS(content, "templates/*.html", "templates/common/*.html")
	if err != nil {
		logger.Error(fmt.Sprintf("Error rendering HTML template: %s", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	data := PageData{
		Title: "Grapnel - FAQ",
	}

	SetCacheHeaders(w, ContentTypeHTML, constants.CacheDuration, "Faq-Html")

	err = tmpl.ExecuteTemplate(w, "faq.html", data)
	if err != nil {
		logger.Error(fmt.Sprintf("Error rendering HTML template: %s", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
