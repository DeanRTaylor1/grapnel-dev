package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DeanRTaylor1/deans-site/logger"
)

type SubscribeRequest struct {
	Email string `json:"email"`
}

func Subscribe(w http.ResponseWriter, r *http.Request, l logger.Logger) {

	var subscribeRequest SubscribeRequest
	err := json.NewDecoder(r.Body).Decode(&subscribeRequest)
	if err != nil {
		ErrorHandler(w, err, http.StatusBadRequest, "Email already exists.")
		return
	}

	email := subscribeRequest.Email

	l.Info(fmt.Sprintf("Email: %s", email))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "success"}`))
}

func ErrorHandler(w http.ResponseWriter, err error, statusCode int, message string) {
	if err != nil {
		errorMessage := map[string]string{"error": err.Error(), "message": message}
		jsonResponse, marshalErr := json.Marshal(errorMessage)
		if marshalErr != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write(jsonResponse)
	}
}
