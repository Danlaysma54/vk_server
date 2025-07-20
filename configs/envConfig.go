package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func getDbURL() string {
	return os.Getenv("DB_URL")
}
func getSourceURL() string {
	return os.Getenv("DB_SOURCE")
}
func getConnString() string {
	return os.Getenv("CONN_STRING")
}

func getJWT_SECRET() string {
	return os.Getenv("JWT_SECRET")
}
