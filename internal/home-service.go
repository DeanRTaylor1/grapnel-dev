package services

import (
	"net/http"

	"github.com/DeanRTaylor1/deans-site/logger"
)

func Home(w http.ResponseWriter, r *http.Request, logger *logger.Logger) {
	logger.Debug("Accessed route: '/'")
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<html><body><h1>Hello, World!</h1></body></html>"))
}
