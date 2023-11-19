package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/DeanRTaylor1/deans-site/logger"
)

func ServeAbout(w http.ResponseWriter, r *http.Request, logger *logger.Logger) {
	logger.Debug("Accessed route: '/faq'")

	tmpl, err := template.ParseFS(content, "templates/*.html")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)

	err = tmpl.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		logger.Error(fmt.Sprintf("Error rendering HTML template: %s", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
