package db

import (
	"github.com/Kenny477/share-img/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}
}

func MigrateDB() {
	DB.AutoMigrate(&models.File{})
}
