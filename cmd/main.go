package main

import (
	rcon "github.com/Random7/go-mc-rcon/cmd/rcon"
	webserver "github.com/Random7/go-mc-rcon/cmd/webserver"
)

func main() {

	webserver.SetupWeb()
	rcon.SetupRcon()

	print("it lives")
}
