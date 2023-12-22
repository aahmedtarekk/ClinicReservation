package initializers

import (
	"fmt"
	"os"
	"toolsProject/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var ActivePool = make(map[string]models.UserPair)

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database")
	}
}

func SyncDB() {
	DB.AutoMigrate(&models.Doctor{})
	DB.AutoMigrate(&models.Patient{})
	DB.AutoMigrate(&models.Slot{})
	DB.AutoMigrate(&models.Appointment{})
}
