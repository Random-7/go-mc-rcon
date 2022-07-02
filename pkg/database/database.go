package database

import (
	"log"

	"github.com/Random7/go-mc-rcon/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConnection struct {
	Dsn string
	Db  *gorm.DB
}

func Connect(connection DbConnection) (*gorm.DB, error) {
	var err error
	connection.Db, err = gorm.Open(mysql.Open(connection.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Error: " + err.Error())
		return nil, err
	}

	logger.PrintInfo("connected to: " + connection.Db.Name())
	return connection.Db, nil
}

func GetFirstRecord(db *gorm.DB) {

}

// connection.db.AutoMigrate(&models.MCServer{})
// connection.db.Create(&models.MCServer{Name: "SamePlanLessDying", Address: "10.0.50.50", Port: "25565", RconPort: "25575", RconPassword: "spldmcrcon2022"})

// var server models.MCServer
// connection.db.First(&server, 1)
// logger.PrintMessage(server.Address)
