package main

import (
	"github.com/Napigo/go-finance-service/configs"
	fr "github.com/Napigo/go-finance-service/frameworks/rest"
	it "github.com/Napigo/go-finance-service/interfaces"
	"github.com/Napigo/go-finance-service/internals/db"
)

func main() {
	// Loads all .env variables before start
	configs.PreloadSettings()

	// Establish connection to Database
	db.ConnectDB()

	// Running the ffirst framework - HTTP Rest Server
	var restServer it.Framework = fr.RestServer{}
	restServer.Run()
}
