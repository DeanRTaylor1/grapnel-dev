package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ENV           string
	IsProduction  bool
	IsTest        bool
	IsDevelopment bool
	BaseUrl       string
	Port          string
	Api_Version   string
}

var Env EnvConfig

func LoadEnv() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory. Error: %s", err.Error())
	}
	projectRoot := filepath.Join(currentDir, "..")

	env := os.Getenv("GO_ENV")
	if env == "" {
		os.Setenv("GO_ENV", "development")
		env = "development"
	}

	fmt.Println(projectRoot)
	if env != "production" {
		envFilePath := filepath.Join(projectRoot, ".env."+env+".local")
		err := godotenv.Load(envFilePath)
		if err != nil {
			log.Fatalf("Error loading env file. Error: %s", err.Error())
		}
	}

	Env = EnvConfig{
		ENV:           getEnv("GO_ENV", "development"),
		IsProduction:  getEnv("GO_ENV", "production") == "production",
		IsTest:        getEnv("GO_ENV", "test") == "test",
		IsDevelopment: getEnv("GO_ENV", "development") == "development",
		BaseUrl:       getEnv("BASE_URL", "http://localhost"),
		Port:          getEnv("PORT", "8080"),
		Api_Version:   getEnv("API_VERSION", "v1"),
	}
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
