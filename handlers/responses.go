package handlers

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func ErrorHandler(w http.ResponseWriter, err error, statusCode int, message string) {
	errorMessage := map[string]string{"message": message}

	if err != nil {
		errorMessage["error"] = err.Error()
	}

	jsonResponse, marshalErr := json.Marshal(errorMessage)
	if marshalErr != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}
