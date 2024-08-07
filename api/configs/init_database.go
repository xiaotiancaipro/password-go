package configs

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"password-go/models"
)

var DB *gorm.DB

func migrateDatabase() {
	DB.AutoMigrate(&models.User{})
}

func InitDB() {
	config := Config().Database
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host,
		config.Username,
		config.Password,
		config.Database,
		config.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	DB = db
	migrateDatabase()
}

func GetDB() *gorm.DB {
	return DB
}
