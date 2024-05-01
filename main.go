package main

import (
	"cars/internal/database"
	"cars/internal/rest"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

// @title           Cars API
// @version         1.0
// @description     This is a testing api for cars

// @host      localhost:8080
// @BasePath  /cars

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	godotenv.Load()
	sqlConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := database.NewPostgresDB(sqlConnectionString)
	if err != nil {
		slog.Debug("Error when connect to database", err)
	}

	rest.StartServer(rest.NewRouter(db), os.Getenv("APP_PORT"))
}
