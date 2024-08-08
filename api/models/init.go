package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"password-go/configs"
	"password-go/utils"
)

var log = utils.LoggerUtil{}.Logger()

func migrateDatabase(db *gorm.DB) {
	tables := map[string]any{
		"user": User{},
	}
	for name, table := range tables {
		err := db.AutoMigrate(&table)
		if err != nil {
			log.Error("Failed to migrate table %s: %v", name, err)
		}
		log.Info("Table %s migration successful", name)
	}

}

func InitDB(config configs.ConfigYaml) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Database.Host,
		config.Database.Username,
		config.Database.Password,
		config.Database.Database,
		config.Database.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("Failed to connect database: %v", err)
	}
	migrateDatabase(db)
	return db
}
