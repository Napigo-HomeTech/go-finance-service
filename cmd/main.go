package main

import (
	log "github.com/Napigo/go-finance-service/pkg/logger"
	"github.com/joho/godotenv"
)

// Loading all env from .env files
func loadEnvironments() {
	err := godotenv.Load()
	if err != nil {
		log.GetLogger().Fatal("Could not load .env files to system")
	}
}

var logger = log.GetLogger()

func main() {
	loadEnvironments()

	logger.Info("oiiii")
	logger.Error("oiiii")
	logger.Warn("Helppp")
	logger.Debug("oiiii")
}
