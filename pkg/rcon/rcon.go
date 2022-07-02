package rcon

import (
	"log"

	mcrcon "github.com/Kelwing/mc-rcon"
	"github.com/Random7/go-mc-rcon/pkg/logger"
)

var rconConnection *mcrcon.MCConn

type ConnectionInfo struct {
	IP       string
	Password string
	Socket   *mcrcon.MCConn
}

func NewRconConnection(auth *ConnectionInfo) error {
	logger.PrintInfo("Setting up Rcon")
	rconConnection = new(mcrcon.MCConn)
	err := rconConnection.Open(auth.IP, auth.Password)
	if err != nil {
		return err
	}

	err = rconConnection.Authenticate()
	if err != nil {
		log.Fatalln("Auth failed", err)
	}
	logger.PrintInfo("Authentication is complete")

	return nil
}

func CloseRconConnection(connection *mcrcon.MCConn) {
	connection.Close()
}

func SendCommand(command string) (string, error) {

	if rconConnection == nil {
		connectionInfo := new(ConnectionInfo)
		connectionInfo.IP = "10.0.50.50:25575"
		connectionInfo.Password = "spldrconmc2022"
		NewRconConnection(connectionInfo)
	}

	resp, err := rconConnection.SendCommand(command)
	if err != nil {
		logger.PrintInfo("Command Failed - " + err.Error())
		return command, err
	}
	logger.PrintInfo("Command Sent - " + resp)

	return "", nil
}
