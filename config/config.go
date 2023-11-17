package config

import (
	"os"

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

	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}

	if env != "production" {
		godotenv.Load(".env." + env + ".local")
	}

	Env = EnvConfig{
		ENV:           getEnv("GO_ENV", "development"),
		IsProduction:  getEnv("GO_ENV", "development") == "production",
		IsTest:        getEnv("GO_ENV", "development") == "test",
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

// func getEnvAsBool(name string, defaultVal bool) bool {
// 	valStr := getEnv(name, "")
// 	if val, err := strconv.ParseBool(valStr); err == nil {
// 		return val
// 	}
// 	return defaultVal
// }

// func getEnvAsInt(name string, defaultVal int) int {
// 	valStr := getEnv(name, "")
// 	if val, err := strconv.Atoi(valStr); err == nil {
// 		return val
// 	}
// 	return defaultVal
// }
