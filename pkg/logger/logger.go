package logger

import "log"

func PrintFatalError(errormessage string, err error) {
	log.Fatalln("ERROR: " + errormessage + " " + err.Error())
}

func PrintMessage(message string) {
	log.Println("INFO: " + message)
}
