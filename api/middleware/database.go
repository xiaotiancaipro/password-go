package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"password-go/configs"
	"password-go/models"
)

func migrateDatabase(db *gorm.DB) {
	tables := map[string]any{
		"user": models.User{},
	}
	for name, table := range tables {
		err := db.AutoMigrate(&table)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to migrate table [%s]: %v", name, err))
			return
		}
		log.Info(fmt.Sprintf("Table [%s] migration successful", name))
	}
}

func InitDB(config configs.ConfigYaml) gin.HandlerFunc {
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
		log.Error(fmt.Sprintf("Failed to connect database: %v", err))
		return nil
	}
	migrateDatabase(db)
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func GetDB(c *gin.Context) (db *gorm.DB, err any) {
	cdb, exists := c.Get("db")
	if !exists {
		db, err = nil, "Failed to get db context"
		return
	}
	gdb, ok := cdb.(*gorm.DB)
	if !ok {
		db, err = nil, "Type assertion failed"
		return
	}
	db, err = gdb, nil
	return
}
