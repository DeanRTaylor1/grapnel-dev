package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/DeanRTaylor1/deans-site/logger"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SubscribeRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func Subscribe(w http.ResponseWriter, r *http.Request, l logger.Logger, client *mongo.Client, v *validator.Validate) {
	var subscribeRequest SubscribeRequest
	err := json.NewDecoder(r.Body).Decode(&subscribeRequest)
	if err != nil {
		ErrorHandler(w, err, http.StatusBadRequest, "Invalid request format.")
		return
	}

	if err := v.Struct(subscribeRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(subscribeRequest.Email)

	collection := client.Database("sysd").Collection("mailing_list")

	var existingEntry SubscribeRequest

	err = collection.FindOne(context.Background(), bson.D{{Key: "email", Value: subscribeRequest.Email}}).Decode(&existingEntry)

	if err == nil {
		ErrorHandler(w, nil, http.StatusConflict, "Email already exists")
		return
	} else if err != mongo.ErrNoDocuments {
		ErrorHandler(w, err, http.StatusInternalServerError, "Error checking existing email.")
		return
	}

	_, err = collection.InsertOne(context.Background(), subscribeRequest)
	if err != nil {
		ErrorHandler(w, err, http.StatusInternalServerError, "Error inserting new email.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "success"}`))
}

func Unsubscribe(w http.ResponseWriter, r *http.Request, l logger.Logger, client *mongo.Client, v *validator.Validate, email string) {
	if email == "" {
		ErrorHandler(w, errors.New("missing email parameter"), http.StatusBadRequest, "Email parameter is required.")
		return
	}

	if err := v.Var(email, "required,email"); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := client.Database("sysd").Collection("mailing_list")

	result, err := collection.DeleteOne(context.Background(), bson.M{"email": email})
	if err != nil {
		ErrorHandler(w, err, http.StatusInternalServerError, "Error deleting email.")
		return
	}

	if result.DeletedCount == 0 {
		ErrorHandler(w, nil, http.StatusNotFound, "Email not found.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "success"}`))
}
