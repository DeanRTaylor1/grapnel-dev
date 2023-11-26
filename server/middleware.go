package server

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/DeanRTaylor1/deans-site/constants"
	"golang.org/x/time/rate"
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

var (
	ipLimiterMap   = make(map[string]*rate.Limiter)
	ipLimiterMutex = &sync.Mutex{}
)

func NewIPRateLimiter(ip string) *rate.Limiter {
	limiter := rate.NewLimiter(rate.Limit(25), 15)
	return limiter
}

func getClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		return strings.Split(xForwardedFor, ",")[0]
	}
	xRealIP := r.Header.Get("X-Real-IP")
	if xRealIP != "" {
		return xRealIP
	}
	return strings.Split(r.RemoteAddr, ":")[0]
}

func limitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := getClientIP(r)

		fmt.Println(clientIP)
		ipLimiterMutex.Lock()
		limiter, exists := ipLimiterMap[clientIP]
		if !exists {
			limiter = NewIPRateLimiter(clientIP)
			ipLimiterMap[clientIP] = limiter
		}
		ipLimiterMutex.Unlock()

		if !limiter.Allow() {
			// Return a 429 Too Many Requests response if the limit is exceeded.
			http.Error(w, "429 Too Many Requests", http.StatusTooManyRequests)
			return
		}
		fmt.Println("checking")

		next.ServeHTTP(w, r)
	})
}

func GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			gw := gzip.NewWriter(w)
			defer gw.Close()
			w.Header().Set("Content-Encoding", "gzip")
			gzippedResponseWriter := gzipResponseWriter{Writer: gw, ResponseWriter: w}
			next.ServeHTTP(gzippedResponseWriter, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (g gzipResponseWriter) Write(b []byte) (int, error) {
	return g.Writer.Write(b)
}
