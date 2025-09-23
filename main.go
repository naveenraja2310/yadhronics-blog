package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	awssettings "yadhronics-blog/aws-settings"
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

	awssettings.InitializeS3Client()
	settings.Log.Info("AWS S3 Client Initialized")

	database.InitDB(config)
	settings.Log.Info("Database Connected")

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
		1,    // maximum size in MB before log is rotated
		3,    // maximum number of old log files to retain
		30,   // maximum number of days to retain old log files
		true, // whether to compress old log files
		"DEBUG",
	)
}
