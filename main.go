package main

import (
	fr "github.com/Napigo/go-finance-service/frameworks/rest"
	it "github.com/Napigo/go-finance-service/interfaces"
	"github.com/Napigo/npgcommon"
)

func main() {
	npgcommon.LoadEnv()

	// Running the ffirst framework - HTTP Rest Server
	var restServer it.Framework = fr.RestServer{}
	restServer.Run()
}
