package main

import (
	rest "github.com/Napigo/go-finance-service/frameworks/rest"
	it "github.com/Napigo/go-finance-service/interfaces"
	"github.com/Napigo/go-finance-service/internals/db"
	"github.com/Napigo/npgc"
)

func main() {
	// Loads all .env variables before start
	npgc.Load(".env")
	// Establish connection to Database
	db.ConnectDB()

	// List of all servers need to initialzied and run
	// such as Rest API, grpc, Websocket etc..
	servers := []it.Framework{
		rest.RestServer{},
	}

	for _, server := range servers {
		server.Run()
	}
}
