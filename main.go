package main

import (
	"cars/internal/database"
	"fmt"
)

func main() {
	// time.NewTimer(time.Second * 1).Stop()
	db, err := database.NewPostgresDB()
	if err != nil {
		fmt.Println(err)
	}
	database.Migrate(db)
}
