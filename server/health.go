package server

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func (s *Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	cpuUsage := runtime.NumCPU()
	memStats := new(runtime.MemStats)
	runtime.ReadMemStats(memStats)
	ramUsage := memStats.HeapAlloc / (1024 * 1024) // in megabytes

	elapsedTime := time.Since(startTime).Milliseconds()

	response := fmt.Sprintf("Server is healthy\nCPU Usage: %d\nRAM Usage: %d MB\nResponse Time: %d ms", cpuUsage, ramUsage, elapsedTime)

	s.Logger.Debug(fmt.Sprintf("Health Check:\n%s", response))

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(response))
	if err != nil {
		s.Logger.Error(fmt.Sprintf("Error writing health check response: %s", err))
	}
}
