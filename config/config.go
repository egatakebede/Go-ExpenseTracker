package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
	JWTSecret string
}

func Load() Config {
	err := godotenv.Load() ; err != nil {
		fmt.Println("No .env file found, using environment variables")
	}
return Config{
        Port:       getEnv("PORT", "9080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "expense_tracker"),
		JWTSecret:  getEnv("JWT_SECRET", "supersecretkey"),

} 


// getEnv reads an environment variable or returns a default if not set
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
