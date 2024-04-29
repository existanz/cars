package main

import (
	"cars/internal/database"
	"cars/internal/rest"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	sqlConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := database.NewPostgresDB(sqlConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	database.Migrate(db.DB)
	// database.FillTestData(db)

	rest.StartServer(rest.NewRouter(), os.Getenv("APP_PORT"))
}
