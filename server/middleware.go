package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/DeanRTaylor1/deans-site/constants"
)

type statusResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func colorLog(marker, method string, path string, statusCode int, duration time.Duration) string {
	ms := float64(duration) / float64(time.Millisecond)
	switch {
	case statusCode >= 500:
		return fmt.Sprintf("[%s]: %s%s%s%s%s%s%s%s", marker, constants.BgRed, method, constants.Reset, constants.Separator, path, constants.Separator, constants.BgRed, fmt.Sprintf("%d", statusCode)+constants.Reset+constants.Separator+fmt.Sprintf("%.3fms", ms))
	case statusCode >= 400:
		return fmt.Sprintf("[%s]: %s%s%s%s%s%s%s%s", marker, constants.BgYellow, method, constants.Reset, constants.Separator, path, constants.Separator, constants.BgYellow, fmt.Sprintf("%d", statusCode)+constants.Reset+constants.Separator+fmt.Sprintf("%.3fms", ms))
	case statusCode >= 300:
		return fmt.Sprintf("[%s]: %s%s%s%s%s%s%s%s", marker, constants.BgCyan, method, constants.Reset, constants.Separator, path, constants.Separator, constants.BgCyan, fmt.Sprintf("%d", statusCode)+constants.Reset+constants.Separator+fmt.Sprintf("%.3fms", ms))
	default:
		return fmt.Sprintf("[%s]: %s%s%s%s%s%s%s%s", marker, constants.BgGreen, method, constants.Reset, constants.Separator, path, constants.Separator, constants.BgGreen, fmt.Sprintf("%d", statusCode)+constants.Reset+constants.Separator+fmt.Sprintf("%.3fms", ms))
	}
}

func ColorLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		srw := &statusResponseWriter{
			ResponseWriter: w,
		}

		next.ServeHTTP(srw, r)

		duration := time.Since(start)

		method := r.Method
		path := r.RequestURI
		statusCode := srw.status

		log.Println(colorLog("Chi", method, path, statusCode, duration))
	})
}
