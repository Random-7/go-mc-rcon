package main

import (
	"log"

	"github.com/Random7/go-mc-rcon/pkg/logger"
	"github.com/Random7/go-mc-rcon/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "go-mc-rcon:Fingas@1437@tcp(10.0.10.5:3306)/go-mc-rcon?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Error: " + err.Error())
	}

	logger.PrintMessage("connected to: " + db.Name())

	db.AutoMigrate(&models.MCServer{})
	db.Create(&models.MCServer{Name: "SamePlanLessDying", Address: "10.0.50.50", Port: "25565", RconPort: "25575", RconPassword: "spldmcrcon2022"})

	var server models.MCServer
	db.First(&server, 1)
	logger.PrintMessage(server.Address)

}
