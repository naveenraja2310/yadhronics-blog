package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := router.Listen(fmt.Sprintf(":%s", config.AppPort)); err != nil {
			settings.Log.Fatal(fmt.Sprintf("Failed to start server: %v", err))
		}
	}()

	<-quit

	settings.Log.Info("Shutting down server...")

}

func initializeConfig() (settings.Configuration, error) {
	return settings.InitConfig()
}

func initializeLogger() {
	settings.InitLogger(
		"yadhronics-blog.log",
		1,
		3,
		30,
		true,
		"DEBUG",
	)
}
