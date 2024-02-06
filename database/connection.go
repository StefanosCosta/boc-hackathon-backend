package database

import (
	"OpenBankingAPI/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root@/openBanking?parseTime=true"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	if err = connection.AutoMigrate(&models.User{}); err != nil {
		return
	}
	if err = connection.AutoMigrate(&models.AuditorAccess{}); err != nil {
		fmt.Println("Failed to auto-migrate Auditor Access Table")
	}

	if err = connection.AutoMigrate(&models.AuditorRequests{}); err != nil {
		fmt.Println("Failed to auto-migrate Auditor Access Table")
	}
	if err = connection.AutoMigrate(&models.Account{}); err != nil {
		fmt.Println("Failed to auto-migrate Accounts Table")
	}
}
