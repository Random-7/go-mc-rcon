package models

import (
	"gorm.io/gorm"
)

type MCServer struct {
	gorm.Model
	Name         string
	Address      string
	Port         string
	RconPort     string
	RconPassword string
}
