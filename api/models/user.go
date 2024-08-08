package models

import "gorm.io/gorm"

type User struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Username string
	Password string
	Email    string `gorm:"unique"`
}

type UserOperate struct{}

func (u UserOperate) IsTableEmpty(db *gorm.DB) bool {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		log.Error("Error querying database")
		return false
	}
	return result.RowsAffected == 0
}

func (u UserOperate) Create(db *gorm.DB, user *User) bool {
	if err := db.Create(user).Error; err != nil {
		log.Error("Error creating user")
		return false
	}
	return true
}
