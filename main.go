package main

import (
	fr "github.com/Napigo/go-finance-service/frameworks/rest"
	it "github.com/Napigo/go-finance-service/interfaces"
	"github.com/Napigo/go-finance-service/pkg/logger"
	"github.com/joho/godotenv"
)

// Loading all env from .env files
func loadEnvironments() {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Could not load .env files to system")
	}
	logger.Info("Environment loaded...")
}

func main() {
	loadEnvironments()

	// Running the ffirst framework - HTTP Rest Server
	var restServer it.Framework = fr.RestServer{}
	restServer.Run()
}
