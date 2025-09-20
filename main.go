package main

import (
	"fmt"
	"os"
	"yadhronics-blog/database"
	"yadhronics-blog/router"
	"yadhronics-blog/settings"
)

func main() {
	config, err := initializeConfig()
	if err != nil {
		fmt.Println("Not able to get config files")
		os.Exit(1)
	}

	initializeLogger()
	settings.Log.Info("Logger Initialized")

	database.InitDB(config)
	router := router.GetRouter()
	if err := router.Listen(fmt.Sprintf(":%d", 8733)); err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}

func initializeConfig() (settings.Configuration, error) {
	return settings.InitConfig()
}

func initializeLogger() {
	settings.InitLogger(
		"yadhronics-blog",
		1,
		3,
		30,
		true,
		"DEBUG",
	)
}
