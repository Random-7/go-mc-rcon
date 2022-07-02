package logger

import "log"

func PrintFatalError(errormessage string, err error) {
	log.Fatalln("FATL ERROR: " + errormessage + " " + err.Error())
}

func PrintInfo(message string) {
	log.Println("INFO: " + message)
}

func PrintErrorMessage(message string) {
	log.Println("ERROR: " + message)
}
