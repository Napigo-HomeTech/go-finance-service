package main

import (
	rest "github.com/Napigo/go-finance-service/frameworks/rest"
	it "github.com/Napigo/go-finance-service/interfaces"
	"github.com/Napigo/go-finance-service/internals/db"
	"github.com/Napigo/npgc"
)

func main() {
	npgc.Load(".env")
	db.ConnectDB()

	servers := []it.Framework{
		rest.RestServer{},
	}

	for _, server := range servers {
		server.Run()
	}
}
