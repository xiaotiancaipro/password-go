package services

import (
	"gorm.io/gorm"
	"password-go/models"
)

type UserService struct{}

func (u UserService) IsTableEmpty(db *gorm.DB) bool {
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		panic("Error querying database:")
	}
	return result.RowsAffected == 0
}
