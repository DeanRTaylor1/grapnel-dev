package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/DeanRTaylor1/deans-site/config"
	"github.com/DeanRTaylor1/deans-site/constants"
	"github.com/DeanRTaylor1/deans-site/logger"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

func ServeContact(w http.ResponseWriter, r *http.Request, logger *logger.Logger) {
	logger.Debug("Accessed route: '/Contact'")

	tmpl, err := template.ParseFS(content, "templates/*.html", "templates/common/*.html")
	if err != nil {
		logger.Error(fmt.Sprintf("Error rendering HTML template: %s", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	data := PageData{
		Title: "Grapnel - Contact Us",
	}

	SetCacheHeaders(w, ContentTypeHTML, constants.CacheDuration, "contact-html")

	err = tmpl.ExecuteTemplate(w, "contact.html", data)
	if err != nil {
		logger.Error(fmt.Sprintf("Error rendering HTML template: %s", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

type ContactRequest struct {
	FirstName      string `json:"first_name" bson:"first_name" validate:"required"`
	LastName       string `json:"last_name" bson:"last_name" validate:"required"`
	CompanyWebsite string `json:"company_website" bson:"company_website" validate:"omitempty,url"`
	Email          string `json:"email" bson:"email" validate:"required,email"`
	Message        string `json:"message" bson:"message" validate:"required"`
}

func PostContact(w http.ResponseWriter, r *http.Request, l logger.Logger, client *mongo.Client, v *validator.Validate, c config.EnvConfig) {
	var contactRequest ContactRequest
	err := json.NewDecoder(r.Body).Decode(&contactRequest)
	if err != nil {
		ErrorHandler(w, err, http.StatusBadRequest, "Invalid request format.")
		return
	}

	if err := v.Struct(contactRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := client.Database(c.Db_Name).Collection(c.Collection_Contact)

	_, err = collection.InsertOne(context.Background(), contactRequest)
	if err != nil {
		ErrorHandler(w, err, http.StatusInternalServerError, "Error inserting new email.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(constants.SuccessResponse))
}
