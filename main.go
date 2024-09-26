package main

import (
	"fmt"
	"task_management/database"
	"task_management/router"
)

func main() {
	database.InitDB()
	router := router.GetRouter()
	if err := router.Listen(fmt.Sprintf(":%d", 8733)); err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}
