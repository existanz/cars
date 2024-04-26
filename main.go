package main

import (
	"cars/internal/database"
	"fmt"
	"time"
)

func main() {
	time.NewTimer(time.Second * 2).Stop()
	_, err := database.NewPostgresDB()
	if err != nil {
		fmt.Println(err)
	}
}
