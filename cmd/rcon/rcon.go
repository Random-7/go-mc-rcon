package rcon

import (
	"log"

	mcrcon "github.com/Kelwing/mc-rcon"
)

func SetupRcon() {
	conn := new(mcrcon.MCConn)
	err := conn.Open("10.0.50.50:25575", "spldmcrcon2022")
	if err != nil {
		log.Fatalln("Open Failed", err)
	}
	defer conn.Close()

	err = conn.Authenticate()
	if err != nil {
		log.Fatalln("Auth failed", err)
	}

	resp, err := conn.SendCommand("tps")
	if err != nil {
		log.Fatalln("Command failed", err)
	}
	log.Println(resp)

}
